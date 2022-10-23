# benedict-foodie-management
A food recording system for the Great Benedict

---
## Functional Requirements
### A database to store food records.
- PostgreSQL
- Data entitiies:
    - Record
        - Food name
        - Description
        - Date of eating
        - Eating quantity
        - Satisfaction score
        - Photo URL
    - Food
        - Food name
        - Nutrition fact
        - Date of purchase
        - Purchase quantity
        - Food type (dry, wet, snack)

### Server: Food-related functions
- Golang Gin Gonic
- To provide CRUD operations on food records.
- A logging module recording the history of editing events.

Optional

- Login module
- Food recommendation

### Client: Calendar page that can record today's food
- ReactNative
- The food description should have
    - Title
    - Description box for text
    - Satisfaction score 1-5 to evaluate how Benedict likes it
    - A photo to show the content
- Use cases:
    1. I can click the calendar box and pop up a form to fill in today's food name, description, score and upload photo.
    2. When submit a record with eaten quantity, the database should update the curent stocking of food.
    3. I can create food through another button to add my current stocking of food.

---
- Evolving data model and interfaces
- Integrating with external APIs
    - Twillio: email notification
    - Create a Slack app which allows you to send and retrieve messages to your server using Slack.
- Scaling capacity
    - how can you evolve your server to support more load?
        - Use Locust and setup 3 different load tests: 
            - one that does only reads
            - one that does only writes
            - one that does a mix of 50% reads and 50% writes
        - How much scale can it tolerate?
        - How can you figure out where you’re spending the most time (hint: try searching for “performance profiling”)?
        - How can you modify your server to support more load (hint: one simple initial strategy might be an in-memory cache, but make sure to think about cache invalidation)?
        - How can you protect the overall stability of the service against too many writes (hint: try searching for “ratelimiting”)?
        - What other techniques could you deploy to scale it up?
        - Is it slow in the server or in the database? How do you know?
    - Write a list of things you’d use to identify which is the case and how you could address it. (Actually making these sorts of fixes might lead you down the path of spending more money on hosting than you want to, so it’s fine if you don’t implement them!)
