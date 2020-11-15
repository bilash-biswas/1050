var
a: int64;
begin
   read(a);
   case(a)of
      61:writeln('Brasilia');
      71:writeln('Salvador');
      11:writeln('Sao Paulo');
      21:writeln('Rio de Janeiro');
      32:writeln('Juiz de Fora');
      19:writeln('Campinas');
      27:writeln('Vitoria');
      31:writeln('Belo Horizonte');
   else
      writeln('DDD nao cadastrado');
   end;
end.
