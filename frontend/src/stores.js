import { writable } from 'svelte/store';
import {main} from '../wailsjs/go/models';


const httpResponse = new main.HttpResponse({});
const currentUrl = writable('');
const currentMethod = writable('get');
const currentParams = writable([{id: 1, "name": "test", "value": "test"}, {id: 2, "name": "test", "value": "test"}]);
const currentResponse = writable(httpResponse);
const currentResponseType = writable('text');

currentUrl.subscribe((value) => {
    // Parse url and get params
    try{
        const url = new URL(value);
        const urlParams = Array.from(url.searchParams);
        let params = [];
        for (let i = 0; i < urlParams.length; i++) {
            params.push({id: i, name:urlParams[i][0], value: urlParams[i][1]});
        }
        currentParams.set(params);
    } catch (e) {
        console.log(e);
    }
    
});

currentResponse.subscribe((value) => {
    if (value.header) {
        const contentType = value.header['Content-Type'];
        if (!contentType) {
            return;
        }
        for (let i = 0; i < contentType.length; i++) {
            if (contentType[i].includes('application/json')) {
                currentResponseType.set('json');
            } else if (contentType[i].includes('application/xml')) {
                currentResponseType.set('xml');
            } else if (contentType[i].includes('text/html')) {
                currentResponseType.set('html');
            } else if (contentType[i].includes('text/plain')) {
                currentResponseType.set('text');
            }
        }
        
    }
});

export { currentUrl, currentResponse, currentParams, currentMethod, currentResponseType };
