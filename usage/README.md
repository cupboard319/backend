# Usage Service

This is the Usage service. 

It tracks usage of the public APIs for each customer by listening for request events from v1 API. As well as using these numbers to power individual tracking it also aggregates these numbers daily to power a ranking/trending system.

It is also the source of truth for usage quotas
- how many requests a customer gets at each subscription tier 
- each customer's individual quota which will be the sum of their subscription quota and any additional quota bundles they may have purchased on top.

It can also be used to track custom events, for example we are using it to store api visits and api calls for registered users. See `SaveEvent` and `ListEvents` endpoints.
