## design doc

> problem: not much objective observability on what happens in life, makes it hard to know certain information like "how much time do I spend on YouTube each week?", or "what are the things correlated with my flare ups of eczema?".

### state

a boolean with a name effectively, this boolean takes on different values over time.

the totality of the observability platform allows different "sources" to provide different states.

examples:

1. is:watching_youtube
2. is:scratching
3. using:certain_laundry_detergent
4. in:certain_location

### span

an active **state** over a certain time interval, this can come with additional metadata for more detailed information.

these spans can be grouped together into a category.

### category

a collection of states under a single umbrella.

### visualization

visualization is very simple, see the illustration below:

![visualization](./visualization.excalidraw)

