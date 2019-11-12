import App from '~/App';
import { reAuthenticate } from '~/utils/auth';

reAuthenticate();

const target = document.querySelector('#app');
const app = new App(
    { target, props: { } }
);

window.app = app;

export default app;