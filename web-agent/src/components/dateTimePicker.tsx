import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DateTimePicker as DTP } from "@mui/x-date-pickers/DateTimePicker";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DemoContainer, DemoItem } from "@mui/x-date-pickers/internals/demo";
import dayjs from "dayjs";

interface DateTimePickerProps {
  handleChange: (selectedDateTime: dayjs.Dayjs | null) => void;
}

export default function DateTimePicker({ handleChange }: DateTimePickerProps) {
  const currentDateTime = dayjs();

  const handleDateTimeChange = (selectedDateTime: dayjs.Dayjs | null) => {
    handleChange(selectedDateTime);
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale="au">
      <DemoContainer components={["DateTimePicker"]}>
        <DemoItem>
          <DTP defaultValue={currentDateTime} onChange={handleDateTimeChange} />
        </DemoItem>
      </DemoContainer>
    </LocalizationProvider>
  );
}
