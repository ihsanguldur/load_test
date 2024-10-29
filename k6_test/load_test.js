import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 200 }, // ramp up
        { duration: '5m', target: 200 }, // stable
        { duration: '30s', target: 0 }, // ramp down
    ]
}

export default function () {
    const url = "http://localhost:8001/load_test";

    const randomPass = Math.floor(1000 + Math.random() * 9000).toString();

    const res = http.post(url, JSON.stringify({ password: randomPass}));
    check(res, {"200": (r) => r.status === 200});
    sleep(1);
}
