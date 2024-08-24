function formatDuration(nanoseconds) {
    const SECONDS = 1e9; // Number of nanoseconds in one second
    const MILLISECONDS = 1e6; // Number of nanoseconds in one millisecond
    const MINUTES = 60 * SECONDS; // Number of nanoseconds in one minute

    if (nanoseconds >= MINUTES) {
        const minutes = nanoseconds / MINUTES;
        return `${minutes.toFixed(2)} min`;
    } else if (nanoseconds >= SECONDS) {
        const seconds = nanoseconds / SECONDS;
        return `${seconds.toFixed(2)} sec`;
    } else if (nanoseconds >= MILLISECONDS) {
        const milliseconds = nanoseconds / MILLISECONDS;
        return `${milliseconds.toFixed(2)} ms`;
    } else {
        return `${nanoseconds} nanoseconds`;
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function copyText(text) {
    try {
        await navigator.clipboard.writeText(text);
    } catch (error) {
        console.error('Failed to copy:', error);
    }
}

export { formatDuration,copyText };