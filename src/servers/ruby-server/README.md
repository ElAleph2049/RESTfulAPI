# Ruby Server API

This Ruby server provides an API to manage users, including operations to get, add, update, and delete users, as well as update hours worked. It uses the `Sinatra` framework to define the API and `sinatra/cross_origin` for enabling Cross-Origin Resource Sharing (CORS).

## Installation Requirements

Before you can run the server, you will need the following installed on your machine:

- [Ruby 2.5+](https://www.ruby-lang.org/en/documentation/installation/)
- [Bundler](https://bundler.io/) for managing Ruby dependencies
- [Git](https://git-scm.com/)
- [curl](https://curl.se/) or any other API testing tool like [Postman](https://www.postman.com/)

### On Mac:

1. **Install Ruby**: You can check if Ruby is installed by running:
   ```bash
   ruby -v
    ```
    If Ruby is not installed, install it using Homebrew:
    ```bash
    brew install ruby
    ```

2. **Install Bundler**: You can check if Bundler is installed by running:
    ```bash
    bundle -v
    ```
    If Bundler is not installed, install it using RubyGems:
    ```bash
    gem install bundler
    ```

3. **Install necessary Ruby gems for the project**:
    ```bash
    bundle install
    ```

4. **Install the required Ruby gems**:
    ```bash
    gem install sinatra
    gem install sinatra-cross_origin
    ```

### On Windows:

1. **Install Ruby**: You can check if Ruby is installed by running:
    ```bash
    ruby -v
    ```
    If Ruby is not installed, download and install Ruby from here - https://www.ruby-lang.org/en/documentation/installation/

2. **Install Bundler**: You can check if Bundler is installed by running:
    ```bash
    bundle -v
    ```
    If Bundler is not installed, install it using RubyGems:
    ```bash
    gem install bundler
    ```

3. **Install necessary Ruby gems for the project**:
    ```bash
    bundle install
    ```

### Running the Server:

1. In the project directory, run the following command to start the server:
    ```bash
    ruby server.rb
    ```

2. The server will start on port 5003 by default. You should see output like:
    ```bash
    == Sinatra (v2.1.0) has taken the stage on 5003 for development with backup from Puma
    Puma starting in single mode...
    * Version 5.5.2 (ruby 2.7.2-p137), codename: Zawgyi
    * Min threads: 0, max threads: 16
    * Environment: development
    * Listening on tcp://localhost:5003
    Use Ctrl-C to stop
    ```

### API Endpoints:

The following endpoints are available:

- GET /users: Get a list of all users.
- GET /users/{id}: Get a user by ID.
- POST /users: Add a new user.
- PUT /users/{id}: Update a user by ID.
- PATCH /users/{id}/hours: Add hours worked for a user.
- DELETE /users/{id}: Delete a user by ID.
- DELETE /users: Delete all users.

