import { Calendar } from "react-native-calendars";

const FoodCalendar = ({ children }) => {
    return <Calendar
        onDayPress={day => {
            console.log('selected day', day);
        }}
        onDayLongPress={day => {
            console.log('selected day', day);
        }}
        // Month format in calendar title. Formatting values: http://arshaw.com/xdate/#Formatting
        monthFormat={'yyyy MM'}
        // Do not show days of other months in month page. Default = false
        hideExtraDays={true}
        // If firstDay=1 week starts from Monday. Note that dayNames and dayNamesShort should still start from Sunday
        firstDay={1}
        // Hide day names. Default = false
        hideDayNames={true}
        // Handler which gets executed when press arrow icon left. It receive a callback can go back month
        onPressArrowLeft={subtractMonth => subtractMonth()}
        // Handler which gets executed when press arrow icon right. It receive a callback can go next month
        onPressArrowRight={addMonth => addMonth()}
        // Disable all touch events for disabled days. can be override with disableTouchEvent in markedDates
        disableAllTouchEventsForDisabledDays={true}
        // Replace default month and year title with custom one. the function receive a date as parameter
        renderHeader={date => {
            /*Return JSX*/
        }}
        // Enable the option to swipe between months. Default = false
        enableSwipeMonths={true}
    />
}

export default FoodCalendar