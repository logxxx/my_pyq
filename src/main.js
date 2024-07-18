import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import "./assets/main.css";
import { ImagePreview } from 'vant';
import { Field } from 'vant';
import { ConfigProvider } from 'vant';
import { Button } from 'vant';
import 'vant/lib/index.css';

createApp(App).use(router).use(ImagePreview).use(Field).use(Button).use(ConfigProvider).mount('#app')
//createApp(App).use(router).use(ImagePreview).mount('#app')
