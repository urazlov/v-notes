<script lang="ts">
  let name: string

  function onEnter(): void {}

  function handleInput(e: Event & { currentTarget: HTMLInputElement }): void {
    name = e.currentTarget.value;
  }
</script>

<div class="ring">
  <i style="--clr: #42aaff"></i>
  <i style="--clr: #0095b6"></i>
  <i style="--clr: #003153"></i>
  <div class="login">
    <h1>Enter your nickname</h1>
    <input autocomplete="off" on:input={handleInput} class="input" id="name" type="text"/>
    <button class="btn" on:click={onEnter}>Enter</button>
  </div>
</div>


<style>

.login {
  z-index: 10;
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 20px;
}

.input {
  width: 250px;
  height: 25px;
  border-radius: 5px;
  border: none;
  padding: 0 10px;
  &:active {
    border: none;
  }
}

.btn {
  width: 100px;
  height: 20px;
  border: 1px solid white;
  border-radius: 5px;
  color: white;
  background-color: transparent;
  cursor: pointer;
}

.ring {
  position: absolute;
  z-index: 1;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 600px;
  height: 500px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.ring i {
  position: absolute;
  inset: 0;
  border: 2px solid #fff;
  transition: 1s;
}
.ring i:nth-child(1) {
  border-radius: 38% 62% 63% 37% / 41% 44% 56% 59%;
  animation: animate 6s linear infinite;
}
.ring i:nth-child(2) {
  border-radius: 41% 44% 56% 59%/38% 62% 63% 37%;
  animation: animate 4s linear infinite;
}
.ring i:nth-child(3) {
  border-radius: 41% 44% 56% 59%/38% 62% 63% 37%;
  animation: animate2 10s linear infinite;
}
.ring:hover i {
  border: 6px solid var(--clr);
  filter: drop-shadow(0 0 20px var(--clr));
}
@keyframes animate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@keyframes animate2 {
  0% {
    transform: rotate(360deg);
  }
  100% {
    transform: rotate(0deg);
  }
}

</style>
