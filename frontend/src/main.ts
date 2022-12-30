import { createApp } from "vue";
import App from "./App.vue";
import "@fortawesome/fontawesome-free/css/all.css";
import "@vueform/slider/themes/default.css";

// Vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { aliases, fa } from "vuetify/iconsets/fa";

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: "dark",
  },
  icons: {
    defaultSet: "fa",
    aliases,
    sets: {
      fa,
    },
  },
});

createApp(App).use(vuetify).mount("#app");
