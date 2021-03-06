require 'net/http'
require 'uri'
require 'json'
require 'openssl'

module Github
  class Client
    def apis
      begin
        uri = URI.parse('https://api.github.com')
        https = Net::HTTP.new(uri.host, uri.port)
        https.use_ssl = true
        res = https.start {
          https.get(uri.request_uri)
        }
        JSON.parse(res.body)
      rescue => e
        p e.message
        nil
      end
    end
  end
end

def github_api()
  client = Github::Client.new
  client.apis
end

if __FILE__ == $0
  p github_api()
end
