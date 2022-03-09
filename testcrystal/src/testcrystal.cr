require "kemal"

get "/ping" do
  "pong"
end

Kemal.run