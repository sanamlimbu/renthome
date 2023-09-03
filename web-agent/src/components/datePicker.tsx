import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DatePicker as DP } from "@mui/x-date-pickers/DatePicker";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import dayjs from "dayjs";

interface DatePickerProps {
  handleChange: (selectedDateTime: dayjs.Dayjs | null) => void;
}
export default function DatePicker({ handleChange }: DatePickerProps) {
  const currentDateTime = dayjs();

  const handleDateTimeChange = (selectedDateTime: dayjs.Dayjs | null) => {
    handleChange(selectedDateTime);
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <DemoContainer components={["DatePicker"]}>
        <DP defaultValue={currentDateTime} onChange={handleDateTimeChange} />
      </DemoContainer>
    </LocalizationProvider>
  );
}
