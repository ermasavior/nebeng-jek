import http from "k6/http";
import { check } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";

//scenario 1 (load test vuser and duration)
export const options = {
  vus: 350, // simulate number of users requesting the API
  duration: "150s", // duration needed for users

  thresholds: {
    http_req_failed: ["rate<0.001"], // the error rate must be lower than 0.1%
    http_req_duration: ["p(90)<5000"], // 90% of requests must complete below 2000ms
    http_req_receiving: ["max<17000"], // max receive request below 17000ms
  },
};

export default function () {
  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization:
        "Bearer eyJhbGciOiJIUzI1NiJ9.eyJkcml2ZXJfaWQiOjJ9.tGaH3DFnWlPnt6eFzSVKoiMshouJ_w8iiTdTHvIEHJQ",
    },
    timeout: "600s",
  };

  const body = {
      is_available: true,
      current_location: {
        longitude: 2,
        latitude: 1
      }
  };

  const set_driver_available = http.patch(
    "http://localhost:9999/v1/drivers/availability",
    JSON.stringify(body),
    params
  );
  check(set_driver_available, {
    "verify success response of set driver available": (set_driver_available) =>
      set_driver_available.status == 200,
  });
}

export function handleSummary(data) {
  return {
    "result/load_vuser_duration.html": htmlReport(data),
    stdout: textSummary(data, { indent: " ", enableColors: true }),
  };
}
