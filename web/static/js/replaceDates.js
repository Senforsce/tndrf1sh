// use strict

export function ReplaceDates () {
    document.querySelector(`a[aria-label="Sunday"]`) && (document.querySelector(`a[aria-label="Sunday"]`).innerHTML = "DIM");
    document.querySelector(`a[aria-label="Monday"]`) && (document.querySelector(`a[aria-label="Monday"]`).innerHTML = "LUN");
    document.querySelector(`a[aria-label="Tuesday"]`) && (document.querySelector(`a[aria-label="Tuesday"]`).innerHTML = "MAR");
    document.querySelector(`a[aria-label="Wednesday"]`) && (document.querySelector(`a[aria-label="Wednesday"]`).innerHTML = "MER");
    document.querySelector(`a[aria-label="Thursday"]`) && (document.querySelector(`a[aria-label="Thursday"]`).innerHTML = "JEU");
    document.querySelector(`a[aria-label="Friday"]`) && (document.querySelector(`a[aria-label="Friday"]`).innerHTML = "VEN");
    document.querySelector(`a[aria-label="Saturday"]`) && (document.querySelector(`a[aria-label="Saturday"]`).innerHTML = "SAM");
    document.querySelectorAll(".fc-list-day-text").forEach(el => {
        let replacement = "";
        switch(el.innerHTML) {
            case "Sunday": {
                (replacement = "Dimanche");
                break;
            }
            case "Monday": {
                (replacement = "Lundi");
                break;
            }
            case "Tuesday": {
                (replacement = "Mardi");
                break;
            }
            case "Wednesday": {
                (replacement = "Mercredi");
                break;
            }
            case "Thursday": {
                (replacement = "Jeudi");
                break;
            }
            case "Friday": {
                (replacement = "Vendredi");
                break;
            }
            case "Saturday": {
                (replacement = "Samedi");
                break;
            }
        }

        el.innerHTML = `${replacement}`;

    })
    document.querySelectorAll(".fc-col-header-cell-cushion").forEach(el => {
        const [day, date] = el.innerHTML.split(" ");
        const dateFr = date ? date.split("/"): ["", ""]
        const [mois, jour] = dateFr
        let replacement = "";
        switch(day) {
            case "Sunday": {
                !date && (replacement = "Dimanche");
                break;
            }
            case "Monday": {
                !date && (replacement = "Lundi");
                break;
            }
            case "Tuesday": {
                !date && (replacement = "Mardi");
                break;
            }
            case "Wednesday": {
                !date && (replacement = "Mercredi");
                break;
            }
            case "Thursday": {
                !date && (replacement = "Jeudi");
                break;
            }
            case "Friday": {
                !date && (replacement = "Vendredi");
                break;
            }
            case "Saturday": {
                !date && (replacement = "Samedi");
                break;
            }


            case "Sun": {
                date && (replacement = `Dim ${jour}/${mois}`);
                break;
            }
            case "Mon": {
                date && (replacement = `Lun ${jour}/${mois}`);
                break;
            }
            case "Tue": {
                date && ( replacement = `Mar ${jour}/${mois}`);
                break;
            }
            case "Wed": {
                date && ( replacement = `Mer ${jour}/${mois}`);
                break;
            }
            case "Thu": {
                date && ( replacement = `Jeu ${jour}/${mois}`);
                break;
            }
            case "Fri": {
                date && ( replacement = `Ven ${jour}/${mois}`);
                break;
            }
            case "Sat": {
                date && ( replacement = `Sam ${jour}/${mois}`);
                break;
            }

            default : {
                replacement = day;
            }
                
        }
        el.innerHTML = `${replacement}`;

        switch(el.innerHTML) {
            case "Dim": {
                !date && (replacement = "Dimanche");
                break;
            }
            case "Lun": {
                !date && (replacement = "Lundi");
                break;

            }
            case "Mar": {
                !date && (replacement = "Mardi");
                break;

            }
            case "Mer": {
                !date && (replacement = "Mercredi");
                break;

            }
            case "jeu": {
                !date && (replacement = "Jeudi");
                break;

            }
            case "Ven": {
                !date && (replacement = "Vendredi");
                break;

            }
            case "Sam": {
                !date && (replacement = "Samedi");
                break;

            }

        }
    })

    document.querySelectorAll(".fc-event-title.fc-sticky").forEach(it => {
        console.log({ HTML: it.innerHTML})
        if(it.innerHTML !== "DemandeEnvoyee") return;
        it.innerHTML = "Demande Envoyée"
    })
}