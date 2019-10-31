import App from '~/App';

const target = document.querySelector('#app');
const app = new App(
    { target, props: { } }
);

window.app = app;

export default app;