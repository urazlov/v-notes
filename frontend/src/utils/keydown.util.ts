export function handleMetaEnterKeydown(event: KeyboardEvent, handler: () => void) {
	if ((event.metaKey || event.ctrlKey) && event.key === 'Enter') {
		event.preventDefault();
		handler();
	}
}

export function handleEnterKeydown(event: KeyboardEvent, handler: () => void) {
	if (event.key === 'Enter') {
		event.preventDefault();
		handler();
	}
}
