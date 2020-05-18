Feature: Test the store proxy and service

Scenario: Store and Retrieve
Given that I invoke Store with key "hello" and value "world"
When I invoke Retrieve with key "hello"
Then I must get back a StoreResponse with key "hello" and value "world"



