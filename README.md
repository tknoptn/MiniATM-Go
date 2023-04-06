# Account Management(ATM) using Go 

## By using this code, a programme that lets the user interact with their bank account . 

### Account Management. The programme has features that let the user make withdrawals, deposits, balance checks, and PIN changes.


creates an account struct with fields for an id, name, balance and pin. It also contains methods for withdrawing, depositing, checking the balance, changing the pin and authenticating the user.

pWhen the user entered a PIN and it checks if it matches the account's PIN. If it does, it displays the account's balance and menu options.

The user can then enter a choice from the menu. 
If it is 1, the user is prompted for an amount to withdraw, and the withdraw method is called. 
If it is 2, the user is prompted for an amount to deposit, and the deposit method is called.
If it is 3, the user is prompted for a new PIN, and the changePin method is called. 
If it is 4, the program exits.

If an invalid choice is made, an error message is displayed.


# Features of the Go used in this 

Structs to store data related to an account (e.g. Account struct)
Methods attached to a struct to perform operations on the data (e.g. 
withdraw(), 
deposit(), 
checkBalance(), 
changePin(), 
authenticate()

Error handling using the error type,
fmt package ,
Conditionals to check for invalid user input , 
Loops to process user input forloop.
