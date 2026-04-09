module SmsSender
  def self.send(phone:, message:)
    # In production, this would use Twilio or similar
    puts "[SMS] Sending to #{phone}: #{message}"
    { status: 'sent', phone: phone }
  end

  def self.send_verification(phone:, code:)
    send(phone: phone, message: "Your ShipFast verification code is: #{code}")
  end

  def self.send_alert(phone:, alert:)
    send(phone: phone, message: "ShipFast Alert: #{alert}")
  end
end
