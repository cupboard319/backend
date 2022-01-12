# Billing Service

The billing service manages the billing of the customer. 

Customers give us money in the following ways
- Pay as you go - add a one off amount of money to your account
- Subscription tiers - Pay us a set amount each month for some level of service. The money is also credited to your balance
- Bundles (Future) - A bundle is a payment that gives you a bundle of usage at some beneficial cost. It could be $10 for an extra 50M requests or similar. 

A billing account is set up which can have multiple users as billing admins to add / remove cards. 

### Responsibility
Usage is recorded against the project where project ID is recorded as the key's owner (TODO change key generation to do this rather than setting the user ID as the key's owner). Usage is tracked by the usage service and balance is tracked by the balance service. The billing service manages the coordination of various activities that results in the customer's payment details being charged. 

### Subscriptions
Subscription tiers are defined in configuration. Each subscription will need an equivalent product in Stripe which will manage the recurring billing, receipts, etc.   

### User flows
Add card - redirect to Stripe checkout to add card
Add funds - use Stripe to charge the saved card
Upgrade/Downgrade subscription to tier X - Set up subscription in Stripe, update quotas etc
