'use strict';
const PAYPAY = require('@paypayopa/paypayopa-sdk-node');
const express = require('express');
const bodyParse = require('body-parser');

const app = express();
app.use(bodyParser.json());

console.log(process.argv);

const paypayClient = new paypay.Client({
    clientId: process.env.PAYPAY_CLIENT_ID,
    clientSecret: process.env.PAYPAY_CLIENT_SECRET,
    merchantId: process.env.PAYPAY_MERCHANT_ID,
    productionMode: false // true for production
});

const { reserv_id, user_id, store_id, menue_id, reserv_time, reserv_cnt, is_recipt, total_amount } = req.body;

// 支払リクエストを作成
const payload = {
    merchantPaymentId: `test_merchant_payment_id_${reserv_id}`,
    amount: {
        amount: total_amount,
        currency: 'JPY'
    },
    requestedAt: new Date().getTime(),
    redirectUrl: 'https://paypay.ne.jp/',
    redirectType: 'WEB_LINK'
};

try {
    const result = await paypayClient.QRCodeCreate(payload);
    if (result.data) {
        res.json({ url: result.data.url });
    } else {
        res.status(500).json({ error: 'Payment creation failed' });
    }
} catch (error) {
    res.status(500).json({ error: error.message });
}
