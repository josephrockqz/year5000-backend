# Year5000 Backend

Logic behind gathering data & storing/retrieving from database

## Project Structure

- /cmd is entry point for the application
- /internal is heart of the application: business logic and database interactions
- - /app point where all dependencies and logic are collected
- - /config initialization of general app configurations
- - /database methods for interacting with database (database layer)
- - /models structure of database tables (database layer)
- - /services entire business logic of the application (business layer)
- - /transport http-server settings, handlers, ports (transportation layer)
- /config contains static configurations of our app related to the process of building the application
- /api documentation for API, schema files, protocol definition files
- /build configuration files for project build
- /deployments contains files related to deployment
- /web specific components for web applications (static web resources, server-side templates)

### Based on this article: https://medium.com/geekculture/how-to-structure-your-project-in-golang-the-backend-developers-guide-31be05c6fdd9
