console.log("Script loaded")

/**
* @param {HTMLElement} alert
*/
function removeAlert(alert) {
	/** @type {HTMLElement} parent */
	alert.classList.add("alert-remove")
	setTimeout(() => {
		alert.remove()
	}, 300)
}


/**
* @param {MouseEvent} ev 
*/
function cancelAlert(ev) {
	/** @type {HTMLElement} parent */
	const parent = ev.target.offsetParent;
	removeAlert(parent)
}

function newAlertCleaningWorker() {
	const perAlertDuration = 10000;
	const mainLoopDuration = 10500;
	let cleaningLoopTimeoutId;
	function stopLoop() {
		console.log("LOOP STOP")
		clearTimeout(cleaningLoopTimeoutId)
	}
	function continueLoop() {
		console.log("LOOP CONTINUE")
		alertCleaningLoop()
	}
	const alertsContainer = document.getElementById("alerts")
	if (!alertsContainer) {
		return
	}
	alertsContainer.addEventListener("mouseenter", stopLoop)
	alertsContainer.addEventListener("mouseleave", continueLoop)

	/**
	 * @param {HTMLElement} toBeRemoved 
	 */
	function scheduleRemoveAlert(toBeRemoved) {
		let countDownProgress = toBeRemoved.getElementsByClassName("alert-countdown")[0]
		countDownProgress.classList.add("active")
		toBeRemovedTimeoutId = setTimeout(() => {
			removeAlert(toBeRemoved)
		}, perAlertDuration)
	}

	console.log("initialized")
	function alertCleaningLoop() {
		cleaningLoopTimeoutId = setTimeout(alertCleaningLoop, mainLoopDuration)
		console.log("LOOP RUN")
		if (alertsContainer.children.length == 0) {
			return
		}
		scheduleRemoveAlert(alertsContainer.children[0])
	}
	alertCleaningLoop()
}

newAlertCleaningWorker()

console.log(htmx.config.useTemplateFragments)
htmx.config.useTemplateFragments = true;
console.log(htmx.config.useTemplateFragments)
