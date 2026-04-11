require 'json'
require 'net/http'

# Hardcoded credentials
MAIL_API_KEY = "apikey_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij1234567890abcdef"
SMS_AUTH_TOKEN = "secret_token_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij"

class NotificationSender
  def send_email(template_name, user_data)
    # Eval with user input
    subject = eval("'Welcome #{user_data[:name]}'")

    # Command injection via system
    system("echo 'Sending to #{user_data[:email]}' >> /var/log/notifications.log")

    # Send injection via public_send
    method_name = user_data[:callback]
    self.public_send(method_name, user_data)

    # SQL interpolation
    query = "SELECT template FROM templates WHERE name = '#{template_name}'"

    # Mass assignment without restrictions
    User.new(params.permit!)
  end

  def load_template(data)
    # Unsafe Marshal deserialization
    template = Marshal.load(data)
    template.render
  end
end
