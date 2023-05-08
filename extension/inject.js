const regex = new RegExp('\\((\\d+), (\\d+)\\)')
const regex2 = new RegExp('focusPokeByName\\(\'(\\D+)\'\\)')

const encounterMenu = document.querySelector("#full-encounter")

const ordinal_suffix_of = (i) => {
    var j = i % 10,
        k = i % 100;
    if (j == 1 && k != 11) {
        return i + "st";
    }
    if (j == 2 && k != 12) {
        return i + "nd";
    }
    if (j == 3 && k != 13) {
        return i + "rd";
    }
    return i + "th";
}

const orderEncounters = (target, table) => {
    for (let j = 0; j < target.children.length; j++) {
        let monBox = target.children.item(j)
        let match = regex2.exec(monBox.getElementsByTagName("img")[0].getAttribute("onclick"))[1]
        let order = 0
        while (order < table.length) {
            if (match === table[order].Pokemon) {
                break
            }
            order++
        }
        monBox.getElementsByTagName("ruby")[0].textContent = ordinal_suffix_of(order+1)
    }
}

const orderAllEncounters = (type, table) => {
    let index = 0
    while (index < encounterMenu.childElementCount) {
        let e = encounterMenu.children.item(index)
        index++
        if (e.tagName == "P" && e.textContent.includes(type)) {
            break
        }
    }
    let target = encounterMenu.children.item(index)

    if (target.classList.contains("day-pool")) {
        orderEncounters(target.firstElementChild, table.Day)

        if (table.Night != null) {
            target = encounterMenu.children.item(index+1)
            orderEncounters(target.firstElementChild, table.Night)
        } else {
            target = encounterMenu.children.item(index+1)
            orderEncounters(target.firstElementChild, table.Day)
        }

        target = encounterMenu.children.item(index+2)
        if (target.classList != null && target.classList.contains("morning-pool")) {
            if (table.Morning != null) {
                orderEncounters(target.firstElementChild, table.Morning)
            } else {
                orderEncounters(target.firstElementChild, table.Day)
            }
        }
    } else if (table.Day != null) {
        orderEncounters(target.firstElementChild, table.Day)
    } else {
        orderEncounters(target.firstElementChild, table)
    }
}

const observer = new MutationObserver((mutationList, observer) => {
    for (const mutation of mutationList) {
        for (const node of mutation.addedNodes) {
            /** @type {Element} */
            let element = node

            if (element.firstElementChild != null && element.firstElementChild.className == "encounter-minimap") {
                let encalc = document.createElement("button")
                encalc.textContent = "Roll Encounter Order"
                encalc.onclick = async () => {
                    let store = await chrome.storage.local.get(["ckkeytrainerId"])

                    let area = element.getElementsByTagName("h3")[0].textContent
                    let url = "https://ckkey.thedankins.biz/api/" + store.ckkeytrainerId + "/rollencounter/" + area.replace(/\s/g, '');
                    let response = await fetch(url)
                    let json = await response.json()
                    if (json.Walking != null) {
                        orderAllEncounters("Walking", json.Walking)
                    }
                    if (json.Surfing != null) {
                        orderAllEncounters("Surfing", json.Surfing)
                    }
                    if (json.Headbutt != null) {
                        orderAllEncounters("Headbutt", json.Headbutt)
                    }
                    if (json.RockSmash != null) {
                        orderAllEncounters("Rock Smash", json.RockSmash)
                    }
                    if (json.Fishing != null) {
                        orderAllEncounters("Old Rod", json.Fishing.Old)
                        orderAllEncounters("Good Rod", json.Fishing.Good)
                        orderAllEncounters("Super Rod", json.Fishing.Super)
                    }
                    if (area === "Route 34") {
                        orderAllEncounters("Special", json.Special[0].Pool)
                    }
                }
                element.append(encalc)
            }

            if (element.classList != null && element.classList.contains("encounter-pool")) {
                for (const menu of element.getElementsByClassName("wild-calc")) {
                    let match = regex.exec(menu.firstElementChild.getAttribute("onclick"))
                    let dvcalc = document.createElement("button")
                    dvcalc.textContent = "DVs"
                    dvcalc.onclick = async () => {
                        let store = await chrome.storage.local.get(["ckkeytrainerId"])
                        let url = "https://ckkey.thedankins.biz/api/" + store.ckkeytrainerId + "/rolldv/" + match[1] + "?level=" + match[2]
                        let response = await fetch(url)
                        let json = await response.json()

                        let stat = document.createElement("button")
                        stat.textContent = "HP: " + json.HealthDV
                        menu.appendChild(stat)
                        stat = document.createElement("button")
                        stat.textContent = "Att: " + json.DVs.Attack
                        menu.appendChild(stat)
                        stat = document.createElement("button")
                        stat.textContent = "Def: " + json.DVs.Defense
                        menu.appendChild(stat)
                        stat = document.createElement("button")
                        stat.textContent = "Spc: " + json.DVs.Special
                        menu.appendChild(stat)
                        stat = document.createElement("button")
                        stat.textContent = "Spe: " + json.DVs.Speed
                        menu.appendChild(stat)
                    }
                    menu.appendChild(dvcalc)
                }
            }
        }
    }
});

observer.observe(encounterMenu, { attributes: false, childList: true, subtree: true });