<script lang="ts">
	import { UpdateNote } from '../wailsjs/go/main/App';
	import Login from './componets/Login.svelte';
	import Note from './componets/Note.svelte';
	import Notes from './componets/Notes.svelte';

	let screen: 'login' | 'note' | 'notes' = 'login';
	let userName: string = '';
	let currentNote: string = '';
	let currentIndex: number;

	function handleLogin(name: string) {
		userName = name;
		screen = 'notes';
	}

	function handleBackButton() {
		screen = 'login';
	}

	function hanleNoteClick(note: string, index: number) {
		currentNote = note;
		currentIndex = index;
		screen = 'note';
	}

	async function handleSaveNote(note: string, index: number): Promise<void> {
		await UpdateNote(userName, note, index);
		screen = 'notes';
	}
</script>

{#if screen === 'login'}
	<Login onEnter={handleLogin} />
{:else if screen === 'notes'}
	<Notes backToLogin={handleBackButton} onNoteClick={hanleNoteClick} name={userName} />
{:else if screen === 'note'}
	<Note text={currentNote} index={currentIndex} saveNote={handleSaveNote} />
{/if}
