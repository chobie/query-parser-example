#include <stdio.h>
#include "query_types.h"
#include "y.tab.h"

int main() {
	query q = parse_query("select from t where foo = '5' ;");
	printf("table name: %s\n", q.f->table);
	printf("where column: %s, value: %s\n", q.w->column_name, q.w->v->svalue);
	close_query(&q);
	return 0;
}
