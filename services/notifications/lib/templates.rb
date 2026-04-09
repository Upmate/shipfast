module Templates
  WELCOME = {
    subject: 'Welcome to ShipFast!',
    body: <<~HTML
      <h1>Welcome aboard!</h1>
      <p>Thanks for signing up for ShipFast. Your account is ready to go.</p>
      <p>Get started by creating your first project.</p>
    HTML
  }.freeze

  PASSWORD_RESET = {
    subject: 'Reset your password',
    body: <<~HTML
      <h1>Password Reset</h1>
      <p>Click the link below to reset your password:</p>
      <p><a href="%{reset_url}">Reset Password</a></p>
      <p>This link expires in 1 hour.</p>
    HTML
  }.freeze

  INVOICE = {
    subject: 'Your ShipFast invoice',
    body: <<~HTML
      <h1>Invoice #%{invoice_id}</h1>
      <p>Amount: %{amount}</p>
      <p>Due date: %{due_date}</p>
      <p><a href="%{invoice_url}">View Invoice</a></p>
    HTML
  }.freeze

  SUBSCRIPTION_RENEWED = {
    subject: 'Subscription renewed',
    body: <<~HTML
      <h1>Subscription Renewed</h1>
      <p>Your %{plan_name} plan has been renewed for another %{interval}.</p>
      <p>Next billing date: %{next_date}</p>
    HTML
  }.freeze
end
