import http from "k6/http";
import { check } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";

export const options = {
  stages: [
    { duration: '1m', target: 100 }, // Ramp-up to 100 virtual users (VUs) in 1 minute
    { duration: '5m', target: 100 }, // Maintain 100 VUs for 5 minutes
    { duration: '1m', target: 0 },  // Ramp-down to 0 VUs in 1 minute
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
        "Bearer xxxxxxxxxxx",
    },
    timeout: "600s",
  };

  const set_driver_available = http.get(
    "http://127.0.0.1:9999/api/rides/v1/riders/ride/1",
    params
  );
  check(set_driver_available, {
    "verify success response of set driver available": (set_driver_available) =>
      set_driver_available.status == 200,
  });
}

export function handleSummary(data) {
  return {
    "result/get_ride_data.html": htmlReport(data),
    stdout: textSummary(data, { indent: " ", enableColors: true }),
  };
}
