require "securerandom"

module Mailer
  def self.send_email(to:, subject:, body:)
    # In production, this would use an SMTP client or SES
    message_id = SecureRandom.uuid
    puts "[Mailer] Sending to=#{to} subject=#{subject} id=#{message_id}"
    message_id
  end
end
