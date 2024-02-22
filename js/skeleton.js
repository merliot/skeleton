import { WebSocketController, ViewMode } from './common.js'

export function run(prefix, url, viewMode) {
	const skeleton = new Skeleton(prefix, url, viewMode)
}

class Skeleton extends WebSocketController {

	open() {
		super.open()
		if (this.state.DeployParams === "") {
			return
		}
		this.img = document.getElementById("led-img")
		this.setupBtn()
		this.showLed()
	}

	handle(msg) {
		switch(msg.Path) {
		case "click":
			this.saveClick(msg)
			break
		}
	}

	setupBtn() {
		if (this.viewMode === ViewMode.ViewFull) {
			this.btn = document.getElementById("led-btn")
			this.btn.onclick = () => { this.click() }
		}
	}

	updateBtn() {
		if (this.viewMode === ViewMode.ViewFull) {
			this.btn.textContent = this.state.Led.State ? "Turn Off LED" : "Turn On LED"
		}
	}

	showLed() {
		this.img.src = this.state.Led.State ? "images/LED-on.png" : "images/LED-off.png"
		this.updateBtn()
	}

	saveClick(msg) {
		var led = this.state.Led
		led.State = msg.State
		this.showLed()
	}

	click() {
		var led = this.state.Led
		led.State = !led.State
		this.webSocket.send(JSON.stringify({Path: "click", State: led.State}))
		this.showLed()
	}
}
