import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    duration: '10s',
    vus: 10,
    rps: 500,
    discardResponseBodies: true,
    noConnectionReuse: true
}
export default function () {
    http.get('http://localhost:8080/test1');
}