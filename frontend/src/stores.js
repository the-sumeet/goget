import { writable } from 'svelte/store';
import {main} from '../wailsjs/go/models';


const httpResponse = new main.HttpResponse({});
const currentUrl = writable('');
const currentMethod = writable('get');
const currentParams = writable([{id: 1, "name": "test", "value": "test"}, {id: 2, "name": "test", "value": "test"}]);
const currentResponse = writable(httpResponse);

export { currentUrl, currentResponse, currentParams, currentMethod };
