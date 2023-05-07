import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    discardResponseBodies: true,
    noConnectionReuse: true,
  
    scenarios: {
        contacts: {
    
          executor: 'per-vu-iterations',
    
          vus: 400,
    
          iterations: 50,
    
        },
    
      },
}
export default function () {
    http.get('http://localhost:8080/test3');
}