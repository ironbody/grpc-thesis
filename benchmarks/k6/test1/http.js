import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    duration: '30s'
}
export default function () {
    http.get('http://localhost:8080/test1');
}