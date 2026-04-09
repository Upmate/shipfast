FROM ruby:3.3-slim
WORKDIR /app
COPY Gemfile Gemfile.lock ./
RUN bundle install --without development test
COPY . .
EXPOSE 8091
CMD ["ruby", "app.rb", "-o", "0.0.0.0"]
