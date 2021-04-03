package connect

import (
	"net/url"
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

var ConnectionUtils = TConnectionUtils{}

/**
 * A set of utility functions to process connection parameters
 */
type TConnectionUtils struct{}

/**
 * Concatinates two options by combining duplicated properties into comma-separated list
 * @param options1 first options to merge
 * @param options2 second options to merge
 * @param keys when define it limits only to specific keys
 */
func (c *TConnectionUtils) Concat(options1 *config.ConfigParams, options2 *config.ConfigParams, keys ...string) *config.ConfigParams {
	options := config.NewConfigParamsFromValue(options1)
	for _, key := range options2.Keys() {
		value1 := options1.GetAsString(key)
		value2 := options2.GetAsString(key)

		if value1 != "" && value2 != "" {
			if len(keys) == 0 || indexOf(keys, key) >= 0 {
				options.SetAsObject(key, value1+","+value2)
			}
		} else if value1 != "" {
			options.SetAsObject(key, value1)
		} else if value2 != "" {
			options.SetAsObject(key, value2)
		}
	}
	return options
}

func (c *TConnectionUtils) concatValues(value1 string, value2 string) string {
	if value1 == "" {
		return value2
	}
	if value2 == "" {
		return value1
	}
	return value1 + "," + value2
}

/**
 * Parses URI into config parameters.
 * The URI shall be in the following form:
 *   protocol://username@password@host1:port1,host2:port2,...?param1=abc&param2=xyz&...
 *
 * @param uri the URI to be parsed
 * @param defaultProtocol a default protocol
 * @param defaultPort a default port
 * @returns a configuration parameters with URI elements
 */
func (c *TConnectionUtils) ParseUri(uri string, defaultProtocol string, defaultPort int) *config.ConfigParams {
	options := config.NewEmptyConfigParams()

	if uri == "" {
		return options
	}

	uri = strings.TrimSpace(uri)

	// Process parameters
	pos := strings.Index(uri, "?")
	if pos > 0 {
		params := uri[pos+1:]
		uri = uri[:pos]

		paramsList := strings.Split(params, "&")
		for _, param := range paramsList {
			pos := strings.Index(param, "=")
			if pos >= 0 {
				key, _ := url.QueryUnescape(param[:pos])
				value, _ := url.QueryUnescape(param[pos+1:])
				options.SetAsObject(key, value)
			} else {
				param, _ = url.QueryUnescape(param)
				options.SetAsObject(param, "")
			}
		}
	}

	// Process protocol
	pos = strings.Index(uri, "://")
	if pos > 0 {
		protocol := uri[:pos]
		uri = uri[pos+3:]
		options.SetAsObject("protocol", protocol)
	} else {
		options.SetAsObject("protocol", defaultProtocol)
	}

	// Process user and password
	pos = strings.Index(uri, "@")
	if pos > 0 {
		userAndPass := uri[:pos]
		uri = uri[pos+1:]

		pos = strings.Index(userAndPass, ":")
		if pos > 0 {
			options.SetAsObject("username", userAndPass[:pos])
			options.SetAsObject("password", userAndPass[pos+1:])
		} else {
			options.SetAsObject("username", userAndPass)
		}
	}

	// Process host and ports
	// options.setAsObject("servers", c.concatValues(options.getAsString("servers"), uri));
	servers := strings.Split(uri, ",")
	for _, server := range servers {
		pos = strings.Index(server, ":")
		if pos > 0 {
			options.SetAsObject("servers", c.concatValues(options.GetAsString("servers"), server))
			options.SetAsObject("host", c.concatValues(options.GetAsString("host"), server[:pos]))
			options.SetAsObject("port", c.concatValues(options.GetAsString("port"), server[pos+1:]))
		} else {
			options.SetAsObject("servers", c.concatValues(options.GetAsString("servers"), server+":"+convert.StringConverter.ToString(defaultPort)))
			options.SetAsObject("host", c.concatValues(options.GetAsString("host"), server))
			options.SetAsObject("port", c.concatValues(options.GetAsString("port"), convert.StringConverter.ToString(defaultPort)))
		}
	}

	return options
}

/**
 * Composes URI from config parameters.
 * The result URI will be in the following form:
 *   protocol://username@password@host1:port1,host2:port2,...?param1=abc&param2=xyz&...
 *
 * @param options configuration parameters
 * @param defaultProtocol a default protocol
 * @param defaultPort a default port
 * @returns a composed URI
 */
func (c *TConnectionUtils) ComposeUri(options *config.ConfigParams, defaultProtocol string, defaultPort int) string {
	builder := ""

	protocol := options.GetAsStringWithDefault("protocol", defaultProtocol)
	if protocol != "" {
		builder = protocol + "://" + builder
	}

	username := options.GetAsNullableString("username")
	if username != nil && *username != "" {
		builder += *username
		password := options.GetAsNullableString("password")
		if password != nil && *password != "" {
			builder += ":" + *password
		}
		builder += "@"
	}

	servers := ""
	defaultPortStr := ""
	if defaultPort > 0 {
		defaultPortStr = convert.StringConverter.ToString(defaultPort)
	}
	hosts := strings.Split(options.GetAsStringWithDefault("host", "???"), ",")
	ports := strings.Split(options.GetAsStringWithDefault("port", defaultPortStr), ",")
	for index := range hosts {
		if len(servers) > 0 {
			servers += ","
		}

		host := hosts[index]
		servers += host

		port := defaultPortStr
		if len(ports) > index && ports[index] != "" {
			port = ports[index]
		}

		if port != "" {
			servers += ":" + port
		}
	}
	builder += servers

	params := ""
	reservedKeys := []string{"protocol", "host", "port", "username", "password", "servers"}
	for _, key := range options.Keys() {
		if indexOf(reservedKeys, key) >= 0 {
			continue
		}

		if len(params) > 0 {
			params += "&"
		}
		params += url.QueryEscape(key)

		value := options.GetAsNullableString(key)
		if value != nil && *value != "" {
			params += "=" + url.QueryEscape(*value)
		}
	}

	if len(params) > 0 {
		builder += "?" + params
	}

	return builder
}

/**
 * Includes specified keys from the config parameters.
 * @param options configuration parameters to be processed.
 * @param keys a list of keys to be included.
 * @returns a processed config parameters.
 */
func (c *TConnectionUtils) Include(options *config.ConfigParams, keys ...string) *config.ConfigParams {
	if len(keys) == 0 {
		return options
	}

	result := config.NewEmptyConfigParams()

	for _, key := range options.Keys() {
		if indexOf(keys, key) >= 0 {
			result.SetAsObject(key, options.GetAsString(key))
		}
	}

	return result
}

/**
 * Excludes specified keys from the config parameters.
 * @param options configuration parameters to be processed.
 * @param keys a list of keys to be excluded.
 * @returns a processed config parameters.
 */
func (c *TConnectionUtils) Exclude(options *config.ConfigParams, keys ...string) *config.ConfigParams {
	if len(keys) == 0 {
		return options
	}

	if options == nil {
		return nil
	}

	values := options.Clone().(*data.StringValueMap)
	if values == nil {
		return nil
	}
	result := config.NewConfigParamsFromValue(values)
	for _, key := range keys {
		result.Remove(key)
	}
	return result
}
