
interface LoginRequestBody {
    Email: string;
    Password: string;
}

interface RegisterRequestBody {
    Email: string;
    Password: string;
}

interface NewBirthRequestChildrenInfo {
    Gender: "male" | "female";
    IsAlive: boolean;
}

interface NewBirthRequestBody {
    Birthday: string;
    Children: NewBirthRequestChildrenInfo[];
    BirthType: "natural" | "caesarean";
}

interface NewBirthResponseBodyChildrenInfo {
    ChildUUID: string;
    BirthUUID: string;
}

interface NewBirthResponseBody {
    ChildrenInfo: NewBirthResponseBodyChildrenInfo[];
    MotherUUID: string;
    Message: string;
}

interface NewCattleRequestBody {
    TagNumber: string;
    Birthday: string;
    Gender: "male" | "female";
    LastInseminationDate?: string;
    LastGiveBirthDate?: string;
    PregnancyStatus?: "pregnant" | "not-pregnant" | "inseminated";
}

interface NewDeathRecordRequestBody {
    CattleUUID: string;
    Date: string;
    Cause: string;
}

interface NewIllnessRecordRequestBody {
    StartDate: string;
    EndDate?: string;
    Name: string;
    AreAntibioticsUsing: boolean;
    BlocksMilking: boolean;
}

interface FailedPregnancyRequestBody {
    StatusUpdateDate: string;
}

interface NewInseminationRecordRequestBody {
    FatherUUID?: string;
    InseminationDate: string;
    InseminationType: "natural" | "artifical";
}

interface NewPregnancyRequestBody {
    StatusUpdateDate: string;
}

interface NewMilkingRecordRequestBody {
    Date: string;
    MilkAmount: number;
}

interface NewWeightRecordRequestBody {
    Date: string;
    Weight: number;
}
