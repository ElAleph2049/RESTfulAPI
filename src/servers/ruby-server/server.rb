require 'sinatra'
require 'json'
require 'sinatra/cross_origin'

# Enable CORS
configure do
  enable :cross_origin
end

# In-memory storage for users
users = []
next_id = 1

# Helper method to find a user by ID
def find_user_by_id(users, id)
  users.find { |user| user[:id] == id }
end

# Get all users
get '/users' do
  content_type :json
  users.to_json
end

# Get a user by ID
get '/users/:id' do
  content_type :json
  user = find_user_by_id(users, params[:id].to_i)
  if user
    user.to_json
  else
    status 404
    { error: 'User not found' }.to_json
  end
end

# Add a new user
post '/users' do
  content_type :json
  data = JSON.parse(request.body.read)
  name = data['name']
  
  if name.nil? || name.strip.empty?
    status 400
    { error: 'Name is required and must be a non-empty string' }.to_json
  else
    new_user = { id: next_id, name: name.strip, hours_worked: 0 }
    users << new_user
    next_id += 1
    status 201
    new_user.to_json
  end
end

# Update a user by ID
put '/users/:id' do
  content_type :json
  user = find_user_by_id(users, params[:id].to_i)

  if user
    data = JSON.parse(request.body.read)
    name = data['name']

    if name && !name.strip.empty?
      user[:name] = name.strip
    end

    user.to_json
  else
    status 404
    { error: 'User not found' }.to_json
  end
end

# Update hours worked for a user
patch '/users/:id' do
  content_type :json
  user = find_user_by_id(users, params[:id].to_i)

  if user
    data = JSON.parse(request.body.read)
    hours_to_add = data['hoursToAdd']

    if hours_to_add.is_a?(Numeric)
      user[:hours_worked] += hours_to_add
      user.to_json
    else
      status 400
      { error: 'Invalid hoursToAdd value' }.to_json
    end
  else
    status 404
    { error: 'User not found' }.to_json
  end
end

# Delete all users
delete '/users' do
  content_type :json
  users.clear
  next_id = 1
  users.to_json
end

# Delete a user by ID
delete '/users/:id' do
  content_type :json
  user = find_user_by_id(users, params[:id].to_i)

  if user
    users.delete(user)
    user.to_json
  else
    status 404
    { error: 'User not found' }.to_json
  end
end
