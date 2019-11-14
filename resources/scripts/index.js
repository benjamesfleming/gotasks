import * as moment from 'moment';
import { navigateTo } from 'svero';
import App from '~/App';
import { reAuthenticate } from '~/utils/auth';

reAuthenticate();
if (window.location.hash == "") {
  navigateTo('/#/');
}

moment.updateLocale('en', {
    relativeTime: {
      future : 'in %s',
      past   : '%s ago',
      s  : function (number, withoutSuffix) {
        return withoutSuffix ? 'now' : 'a few seconds';
      },
      m  : '1m',
      mm : '%dm',
      h  : '1h',
      hh : '%dh',
      d  : '1d',
      dd : '%dd',
      M  : '1mth',
      MM : '%dmth',
      y  : '1y',
      yy : '%dy'
    }
});

const target = document.querySelector('#app');
const app = new App(
    { target, props: { } }
);

window.app = app;

export default app;