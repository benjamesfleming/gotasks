<script>
/**
 *  @see https://codepen.io/daniandl/pen/OgbXzK 
 */
</script>

<style lang="postcss">
.label__checkbox {
  display: none;
}
.label__check {
  border-radius: 50%;
  border: 5px solid rgba(0, 0, 0, 0.1);
  background: white;
  vertical-align: middle;
  width: 2em;
  height: 2em;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: border 0.3s ease;
}
.label__check i.icon {
  opacity: 0.2;
  font-size: calc(1rem + 1vw);
  color: transparent;
  transition: opacity 0.3s 0.1s ease;
  -webkit-text-stroke: 3px rgba(0, 0, 0, 0.5);
}
.label__check:hover {
  border: 5px solid rgba(0, 0, 0, 0.2);
}
.label__checkbox:checked + .label__text .label__check {
  -webkit-animation: check 0.5s cubic-bezier(0.895, 0.03, 0.685, 0.22) forwards;
          animation: check 0.5s cubic-bezier(0.895, 0.03, 0.685, 0.22) forwards;
}
.label__checkbox:checked + .label__text .label__check .icon {
  opacity: 1;
  -webkit-transform: scale(0);
          transform: scale(0);
  color: white;
  -webkit-text-stroke: 0;
  -webkit-animation: icon 0.3s cubic-bezier(1, 0.008, 0.565, 1.65) 0.1s 1 forwards;
          animation: icon 0.3s cubic-bezier(1, 0.008, 0.565, 1.65) 0.1s 1 forwards;
}
@-webkit-keyframes icon {
  from {
    opacity: 0;
    -webkit-transform: scale(0.3);
            transform: scale(0.3);
  }
  to {
    opacity: 1;
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}
@keyframes icon {
  from {
    opacity: 0;
    -webkit-transform: scale(0.3);
            transform: scale(0.3);
  }
  to {
    opacity: 1;
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}
@-webkit-keyframes check {
  0% {
    width: 1.5em;
    height: 1.5em;
    border-width: 5px;
  }
  10% {
    width: 1.5em;
    height: 1.5em;
    opacity: 0.1;
    background: rgba(0, 0, 0, 0.2);
    border-width: 15px;
  }
  12% {
    width: 1.5em;
    height: 1.5em;
    opacity: 0.4;
    background: rgba(0, 0, 0, 0.1);
    border-width: 0;
  }
  50% {
    width: 2em;
    height: 2em;
    background: #00d478;
    border: 0;
    opacity: 0.6;
  }
  100% {
    width: 2em;
    height: 2em;
    background: #00d478;
    border: 0;
    opacity: 1;
  }
}
@keyframes check {
  0% {
    width: 1.5em;
    height: 1.5em;
    border-width: 5px;
  }
  10% {
    width: 1.5em;
    height: 1.5em;
    opacity: 0.1;
    background: rgba(0, 0, 0, 0.2);
    border-width: 15px;
  }
  12% {
    width: 1.5em;
    height: 1.5em;
    opacity: 0.4;
    background: rgba(0, 0, 0, 0.1);
    border-width: 0;
  }
  50% {
    width: 2em;
    height: 2em;
    background: #00d478;
    border: 0;
    opacity: 0.6;
  }
  100% {
    width: 2em;
    height: 2em;
    background: #00d478;
    border: 0;
    opacity: 1;
  }
}
</style>

<div class="h-8 w-full flex items-center justify-center">
    <label class="label">
        <input  class="label__checkbox" type="checkbox"/>
        <span class="label__text">
            <span class="label__check">
                <i class="fas fa-check"></i>
            </span>
        </span>
    </label>
</div>