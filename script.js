import http from 'k6/http';
import { sleep } from 'k6';

  
export const options = {
  vus: 10,
  duration: '60m',
};

export default function () {
  const url = 'http://payment-payment.apps.hetzner.calopezb.com/api/pay';
  
  const payload = JSON.stringify({
    amount: 100, 
    currency: 'USD'
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
  sleep(1)
}
