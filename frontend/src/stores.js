import { writable } from 'svelte/store';
import {main} from '../wailsjs/go/models';


const httpResponse = new main.HttpResponse({});
const currentUrl = writable('');
const currentMethod = writable('get');
const currentParams = writable({});
const currentResponse = writable(httpResponse);

export { currentUrl, currentResponse, currentMethod };
