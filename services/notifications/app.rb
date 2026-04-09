require "sinatra"
require "json"
require_relative "lib/mailer"

set :port, 8091

get "/health" do
  content_type :json
  { status: "ok", service: "shipfast-notifications" }.to_json
end

post "/notify/email" do
  content_type :json
  body = JSON.parse(request.body.read)

  to = body["to"]
  subject = body["subject"]
  message = body["message"]

  result = Mailer.send_email(to: to, subject: subject, body: message)
  { status: "sent", message_id: result }.to_json
end

post "/notify/sms" do
  content_type :json
  body = JSON.parse(request.body.read)

  { status: "queued", to: body["phone"] }.to_json
end
