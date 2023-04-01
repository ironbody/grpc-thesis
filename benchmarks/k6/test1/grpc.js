import grpc from "k6/net/grpc";

/*
Throughput: Sends 5 gRPC requests to the server using k6's grpc.invoke() function 
and measures the RPS that the server can handle.
*/

export let options_throughput = {
    vus: 10,
    duration: "30s"
};

export default function () {
    let client = new grpc.Client();
    client.load(["./proto/user.proto"]);
    client.connect("localhost:9000", { plaintext: true });

    let requests = [
        { user_id: 1 },
        { user_id: 2 },
        { user_id: 3 },
        { user_id: 4 },
        { user_id: 5 }
    ];

    let responses = client.invoke("UserService/GetUser", requests);

    let rps = requests.length / 30;
    console.log(`gRPC Requests per Second (RPS): ${rps}`);
}


// Latency: Sends a single GET request to the server and measures the response time in milliseconds.
export let options_latency = {
    vus: 10,
    duration: "30s"
};

export default function () {
    let client = new grpc.Client();
    client.load(["./proto/user.proto"]);
    client.connect("localhost:9000", { plaintext: true });

    let request = { user_id: 1 };

    let response = client.invoke("UserService/GetUser", request);
    console.log(`gRPC Response Time (ms): ${response.meta_data.latency}`);
}

// Scalability: Sends requests to the server using 100 VUs and measures the RPS and response time for each VU.
export let options_scalability = {
    vus: 100,
    duration: "30s"
};

export default function () {
    let client = new grpc.Client();
    client.load(["./proto/user.proto"]);
    client.connect("localhost:9000", { plaintext: true });

    let request = { user_id: 1 };

    let response = client.invoke("UserService/GetUser", request);
    let rps = requests.length / 30;
    console.log(`gRPC Requests per Second (RPS): ${rps}`);
    console.log(`gRPC Response Time (ms): ${response.meta_data.latency}`);
}