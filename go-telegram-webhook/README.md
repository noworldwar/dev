To hook the api site:
./ngrok http 3000
curl -F "url=<your_api_url>"  https://api.telegram.org/bot<your_api_token>/setWebhook
Example: curl -F "url=https://3c08-122-3-23-211.ngrok.io"  https://api.telegram.org/bot2066047122:AAGMxupBSwAvpTUpvnIgBrQA_k9cYtT_OzQ/setWebhook