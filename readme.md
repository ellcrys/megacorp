### MegaCorp - Employee Management and Salary Disbursement Contract

This is a sample contract for managing and disbursing salaries of 
a hypothetical company "MegaCorp" that pays its employees in its
own digital currency known as "MegaCoin". This contract based on version 0.5.0 of the Ellcrys platform and is intended for the purpose of illustrating the structure of an Ellcrys contract, state storage and access. 

This contract provides the following functions:

- **create-account** Creates a MegaCorp account. This is required to hold MegaCoin. In MegaCorp's world, anyone who wishes to use MegaCoin will need this account.

- **create-employee**: Create/Add a MegaCorp employee. The created employees will be paid according to their set position (ceo, coo etc). Requires a MegaCorp account.

- **get-account**: Retrieve information about a MegaCorp account.

- **get-total-supply**: Retrieve the current total of un-issued MegaCoins.


**NOTE:**

For simplicity, this contract does not handle authentication/authorization during function execution. 