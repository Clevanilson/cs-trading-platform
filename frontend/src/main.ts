import { createApp } from "vue";

import App from "./App.vue";
import "./assets/tokens/colors.css";
import "./assets/tokens/typography.css";
import "./assets/tokens/sizes.css";
import "./assets/main.css";
import { HttpClientFetchAdapter } from "./http/http-client-fetch-adapter.ts";
import { AccountHttpGateway } from "./gateway/account-http-gateway.ts";

const app = createApp(App);
const httpClient = new HttpClientFetchAdapter();
const accountGateway = new AccountHttpGateway(httpClient);
app.provide("accountGateway", accountGateway);

app.mount("#app");
