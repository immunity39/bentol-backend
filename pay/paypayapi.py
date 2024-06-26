from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

@app.route('/pay', methods=['POST'])
def pay():
    data = request.json
    user_id = data['user_id']
    store_id = data['store_id']
    menue_id = data['menue_id']
    reserv_time = data['reserv_time']
    reserv_cnt = data['reserv_cnt']
    is_recipt = data['is_recipt']
    total_amount = data['total_amount']

    # PayPay APIへのリクエスト
    paypay_api_url = "https://api.paypay.ne.jp/v2/payments"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_PAYPAY_API_KEY"
    }
    payload = {
        "merchantPaymentId": f"{user_id}-{store_id}-{menue_id}",
        "amount": {
            "amount": total_amount,
            "currency": "JPY"
        },
        "orderDescription": "Bento reservation payment",
        "orderItems": [
            {
                "name": "Bento",
                "quantity": reserv_cnt,
                "unit_price": {
                    "amount": total_amount / reserv_cnt,
                    "currency": "JPY"
                }
            }
        ],
        "userAgent": request.headers.get('User-Agent')
    }

    response = requests.post(paypay_api_url, headers=headers, json=payload)
    
    if response.status_code == 200:
        return jsonify({"message": "Payment successful"})
    else:
        return jsonify({"error": "Payment failed"}), response.status_code

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
