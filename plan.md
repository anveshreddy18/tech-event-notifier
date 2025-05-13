Plan: 

- So we basically have the root, fetch and notify commands. We should be writing the interfaces for these functionalities and let the concrete types implement them. Keep the interfaces to a minimum.
- Once that is done for both the commands, these should be used in the command handlers and also in the main function to either spin them up using cli or just running the application binary for full functionality.
[Need to see whether the above is the recommended way of doing it, if yes, why yes. If not, why not, then decide how you want to approach this.]
- We only want to implement the fetching, processing and notifying functionalities. Atleast for the initial version, I don't want to store the data in a database. We can do that later.


Contents from main.go

// in the flow, the first thing to do is to call the appropriate event aggregator platform.
// So our binary should basically have commands such as fetch, notify and each of these will have supporting flags such as the name of the city, date, distance, etc.
// The fetch command will be used to fetch the events from the aggregator platform.
// The notify command will be used to notify the user about the events using the mentioned communication channel.
// TODO: @anveshreddy18 -- In future, we should add flags such as filtering with distance, filtering based on date, etc.


First: Read this blog about SOLID: https://blog.stackademic.com/mastering-solid-principles-with-go-examples-71db32b8c990
REFER TO THE WRAPPER AND BACKUP RESTORE WAY OF PASSING THE FLAGS, DEFINING THE INTERFACES, STRUCTS, ETC. ALSO EXPERIMENT WITH YOUR OWN WAY OF DOING IT. LET'S SEE WHAT WORKS BEST.
EMBIBE THE BIAS FOR ACTION AND THE BIAS FOR DECISION MAKING. F**K AROUND AND FIND OUT IS THE WAY TO GO.

FIRST TRY TO FUCK IT ALL UP. THEN UNDERSTAND WHETHER SOLID PRINCIPLES ARE WORTH IT, IF YES, THEN WE NOW KNOW WHY

Also IMPORTANT: Do proper error handling and pass context properly