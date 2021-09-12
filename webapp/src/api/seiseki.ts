export interface SeisekiJSON {
  ID: number;
  UserID: number;
  Seiseki: InnerSeiseki;
  CreatedAt: Date;
  UpdatedAt: Date;
}

export interface Seiseki extends InnerSeiseki {
  id: number;
  UserID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
}

export interface InnerSeiseki {
  SubjectName: string;
  TeacherName: string;
  SubjectDistinction: string;
  Credit: number;
  Grade: string;
  Score: number;
  GP: number;
  Year: number;
  Date: Date;
}
