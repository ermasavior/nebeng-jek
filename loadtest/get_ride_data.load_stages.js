import http from "k6/http";
import { check } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";

//scenario 2 (load test based on stages)
export const options = {
  stages: [
    { duration: "0.3m", target: 50 },
    { duration: "0.4m", target: 100 },
    { duration: "0.6m", target: 150 },
    { duration: "1m", target: 200 },
    { duration: "1.4m", target: 0 },
  ],
  thresholds: {
    http_req_failed: ["rate<0.001"], // the error rate must be lower than 0.1%
    http_req_duration: ["p(90)<1800"], // 90% of requests must complete below 1800ms
    http_req_receiving: ["max<6000"], // max receive request below 6000ms
  },
};

export default function () {
  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiJ9.eyJyaWRlcl9pZCI6MX0.zdqgIG9PivvL4CiiDC7NrlGGDz6OyQ478KlVE2YmjKE",
    },
    timeout: "600s",
  };

  const set_driver_available = http.get(
    "http://localhost:9999/v1/riders/ride/1",
    params
  );
  check(set_driver_available, {
    "verify success response of set driver available": (set_driver_available) =>
      set_driver_available.status == 200,
  });
}

export function handleSummary(data) {
  return {
    "result/load_stages/get_ride_data.result.html": htmlReport(data),
    stdout: textSummary(data, { indent: " ", enableColors: true }),
  };
}
