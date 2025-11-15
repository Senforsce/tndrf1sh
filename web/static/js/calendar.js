// use strict

export function RenderCalendar(eventsData) {
    const calendarEl = document.getElementById('calendar');
    const calendar = new Calendar(calendarEl, {
        plugins: [index, daygrid_index, timegrid_index, list_index],
        headerToolbar: {
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek,timeGridDay,listWeek'
        },
        initialDate: new Date().toISOString().split('T')[0],
        navLinks: true,
        // can click day/week names to navigate views
        editable: true,
        dayMaxEvents: true,
        // allow "more" link when too many events
        events: eventsData.concat([
        {
            groupId: '999',
            title: 'Repeating Event',
            start: '2025-05-09T16:00:00'
        },
        {
            groupId: '999',
            title: 'Repeating Event',
            start: '2025-05-16T16:00:00'
        }

        ])
    });
    calendar.render();

    return calendarEl;
}

export function GetCalendarData() {
    const data = document.querySelectorAll(".o8_data");
    const lastdata = data.length-1; //always get the last most recent data
    const topropagate = JSON.parse(data[lastdata].dataset.propagate);
    const po = JSON.parse(data[lastdata].dataset.po);
    const bookings = JSON.parse(data[lastdata].dataset.booking);

    const cr = JSON.parse(data[lastdata].dataset.cr);
    let units = []
    cr.forEach(it => {
        if (it.Units) {
            units =  units.concat(it.Units.map(ji => {
                return {
                    title: (it.BelongsToService === "http://senforsce.com/o8/brain/coach-mj/ServiceProgrammePersonnalise" ? "Programme Persos" : "Coaching Individuel" ),
                    start: ji.CreditBookingAgreedTime.replace(",", "")
                }
            }))
        }
    })
    let evtsWrapper = topropagate.map(item => {
        return {
            title: item.MoveName + "PR " + item.PRValue + item.PRUnit ,
            start: item.PRDate,
        }
    }).concat(po.map(( objective ) => {
        return objective.ObjectiveKPIList.map((listitem) => {
            return {
                title: objective.ObjectiveKPIName + ': ' + listitem.ObjectiveKPIValue,
                start: listitem.ObjectiveKPIDate.split("T")[0]
            }

        })[0]
    }))
    .concat(
        units
    ).concat(
        bookings.map((booking) => {
            return {
                title: `${booking.BookingStatusString}`,
                start: booking.BookingRequestDate
            }
        })
    )

    return evtsWrapper.filter(function( element ) {
        return element !== undefined;
    });
}