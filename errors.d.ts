type ErrorResponseMessages =
  | "UNAUTHORIZED"
  | "SERVER_ERROR"
  | "SMTP_ERROR"
  | "RATE_LIMIT_EXCEEDED"
  | "BAD_REQUEST"
  | "WRONG_CREDENTIALS"
  | "EMAIL_NOT_VERIFIED"
  | "EMAIL_ALREADY_IN_USE"
  | "TAG_NUMBER_ALREADY_IN_USE"
  | "BIRTHDAY_PARSING_FAILED"
  | "SEND_LAST_INSEMINATION_DATE"
  | "LAST_INSEMINATION_DATE_PARSING_FAILED"
  | "LAST_GIVE_BIRTH_DATE_PARSING_FAILED"
  | "SEND_PREGNANCY_STATUS"
  | "CATTLE_NOT_FOUND_OR_ALREADY_PREGNANT"
  | "CATTLE_NOT_FOUND"
  | "NOT_FOUND"
  | "INSEMINATION_DATE_PARSING_FAILED"
  | "STATUS_DATE_PARSING_FAILED"
  | "CATTLE_NOT_FOUND_OR_NOT_INSEMINATED"
  | "INSEMINATION_NOT_FOUND_OR_NOT_UNCERTAIN"
  | "CATTLE_NOT_PREGNANT"
  | "CATTLE_NOT_FOUND_OR_NOT_PREGNANT"
  | "CATTLE_NOT_FOUND_OR_NOT_FEMALE"
  | "INSEMINATION_NOT_FOUND"
  | "INSEMINATION_ALREADY_DONE_OR_FAILED"
  | "DEATH_DATE_PARSING_FAILED"
  | "START_DATE_PARSING_FAILED"
  | "END_DATE_PARSING_FAILED"
  | "MILKING_DATE_PARSING_FAILED"
  | "WEIGHT_DATE_PARSING_FAILED"
