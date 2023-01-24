# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> Component definitions for Pip.Services in Go Changelog

## <a name="1.3.2"></a> 1.3.2 (2023-01-12) 

## Bug fixes
* Fixed MemoryCredentialStore reading by keys
* Update dependencies and docs


## <a name="1.3.0"></a> 1.3.0 (2021-10-24) 
Added state management components
### Features
* **state** Added IStateStore interface and StateValue struct
* **state** Added NullStateStore struct
* **state** Added MemoryStateStore struct
* **state** Added DefaultStateStoreFactory struct
* **config** Extended IConfigReader interface (added *AddChangeListener* and *RemoveChangeListener* methods)

## <a name="1.2.0"></a> 1.2.0 (2021-04-23) 

### Features
* Added trace package with Trace components

## <a name="1.1.0"></a> 1.1.0 (2021-04-03) 

### Features
* **connect** Added CompositeConnectionResolver class
* **connect** Added ConnectionUtils class
* **config** Replaced 3rd party Handlebars templating engine with Mustache templates from expressions module

## <a name="1.0.7"></a> 1.0.7 (2021-03-22)

Factory.Register now accepts function with locator parameter

## <a name="1.0.6"></a> 1.0.6 (2021-03-05)

Update ICache, MemoryCache and NullCache

## <a name="1.0.5"></a> 1.0.5 (2020-12-11)

Update dependencies

## <a name="1.0.0"></a> 1.0.0 (2018-06-27)

Initial public release

### Features
* **auth** Authentication and authorization components
* **cache** Memory cache components
* **build** Component factories framework
* **config** Configuration components
* **connect** Connection components
* **count** Performance counters components
* **info** Information components
* **lock** Lock components
* **log** Logging components