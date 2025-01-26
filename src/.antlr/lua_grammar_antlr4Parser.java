// Generated from d:/Programs/Programmed/fluxlang/antlr4/FLUXASSEMBLY/src/lua_grammar_antlr4.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class lua_grammar_antlr4Parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		SEPARATOR=25, NEWLINE=26, KW_IN=27, KW_PRINT=28, KW_AND=29, KW_BREAK=30, 
		KW_DO=31, KW_ELSE=32, KW_ELSEIF=33, KW_END=34, KW_FALSE=35, KW_FOR=36, 
		KW_GOTO=37, KW_IF=38, KW_NIL=39, KW_NOT=40, KW_OR=41, KW_REPEAT=42, KW_RETURN=43, 
		KW_THEN=44, KW_TRUE=45, KW_UNTIL=46, KW_WHILE=47, KW_LOCAL=48, KW_FUNCTION=49, 
		NUMBER=50, STRING=51, LETTER=52, DIGIT=53, WS=54;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_control_statement = 2, RULE_statement_terminator = 3, 
		RULE_assignment = 4, RULE_expression = 5, RULE_operators = 6, RULE_comparison_operator = 7, 
		RULE_arith_operator = 8, RULE_logical_operator = 9, RULE_comment = 10, 
		RULE_literal = 11, RULE_function_call = 12, RULE_table_insert = 13, RULE_function_declaration = 14, 
		RULE_block = 15, RULE_if_statement = 16, RULE_for_statement = 17, RULE_while_statement = 18, 
		RULE_table = 19, RULE_field = 20, RULE_table_access = 21, RULE_multi_line_comment = 22, 
		RULE_single_line_comment = 23, RULE_print_statement = 24, RULE_identifier = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "control_statement", "statement_terminator", 
			"assignment", "expression", "operators", "comparison_operator", "arith_operator", 
			"logical_operator", "comment", "literal", "function_call", "table_insert", 
			"function_declaration", "block", "if_statement", "for_statement", "while_statement", 
			"table", "field", "table_access", "multi_line_comment", "single_line_comment", 
			"print_statement", "identifier"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'('", "')'", "'>'", "'<'", "'>='", "'=='", "'<='", "'~='", 
			"'*'", "'/'", "'+'", "'-'", "','", "'table.insert'", "'{'", "'}'", "'['", 
			"']'", "'--[['", "'--]]'", "']]'", "'--'", "'_'", null, null, "'in'", 
			"'print'", "'and'", "'break'", "'do'", "'else'", "'elseif'", "'end'", 
			"'false'", "'for'", "'goto'", "'if'", "'nil'", "'not'", "'or'", "'repeat'", 
			"'return'", "'then'", "'true'", "'until'", "'while'", "'local'", "'function'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "SEPARATOR", "NEWLINE", "KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", 
			"KW_DO", "KW_ELSE", "KW_ELSEIF", "KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", 
			"KW_IF", "KW_NIL", "KW_NOT", "KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", 
			"KW_TRUE", "KW_UNTIL", "KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "NUMBER", 
			"STRING", "LETTER", "DIGIT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "lua_grammar_antlr4.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public lua_grammar_antlr4Parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<CommentContext> comment() {
			return getRuleContexts(CommentContext.class);
		}
		public CommentContext comment(int i) {
			return getRuleContext(CommentContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8902574129381380L) != 0)) {
				{
				setState(54);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__1:
				case T__14:
				case T__15:
				case KW_PRINT:
				case KW_FALSE:
				case KW_FOR:
				case KW_IF:
				case KW_NIL:
				case KW_TRUE:
				case KW_WHILE:
				case KW_LOCAL:
				case KW_FUNCTION:
				case NUMBER:
				case STRING:
				case LETTER:
					{
					setState(52);
					statement();
					}
					break;
				case T__19:
				case T__22:
					{
					setState(53);
					comment();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Control_statementContext control_statement() {
			return getRuleContext(Control_statementContext.class,0);
		}
		public Function_declarationContext function_declaration() {
			return getRuleContext(Function_declarationContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(64);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(61);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(62);
				function_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(63);
				function_call();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Control_statementContext extends ParserRuleContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public For_statementContext for_statement() {
			return getRuleContext(For_statementContext.class,0);
		}
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public Control_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_statement; }
	}

	public final Control_statementContext control_statement() throws RecognitionException {
		Control_statementContext _localctx = new Control_statementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_control_statement);
		try {
			setState(69);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KW_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(66);
				if_statement();
				}
				break;
			case KW_FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(67);
				for_statement();
				}
				break;
			case KW_WHILE:
				enterOuterAlt(_localctx, 3);
				{
				setState(68);
				while_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Statement_terminatorContext extends ParserRuleContext {
		public TerminalNode SEPARATOR() { return getToken(lua_grammar_antlr4Parser.SEPARATOR, 0); }
		public Statement_terminatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement_terminator; }
	}

	public final Statement_terminatorContext statement_terminator() throws RecognitionException {
		Statement_terminatorContext _localctx = new Statement_terminatorContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_statement_terminator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(SEPARATOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public TerminalNode KW_LOCAL() { return getToken(lua_grammar_antlr4Parser.KW_LOCAL, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(73);
				match(KW_LOCAL);
				}
			}

			setState(76);
			identifier();
			setState(77);
			match(T__0);
			setState(78);
			expression(0);
			setState(79);
			statement_terminator();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public TableContext table() {
			return getRuleContext(TableContext.class,0);
		}
		public Table_accessContext table_access() {
			return getRuleContext(Table_accessContext.class,0);
		}
		public OperatorsContext operators() {
			return getRuleContext(OperatorsContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(82);
				literal();
				}
				break;
			case 2:
				{
				setState(83);
				identifier();
				}
				break;
			case 3:
				{
				setState(84);
				match(T__1);
				setState(85);
				expression(0);
				setState(86);
				match(T__2);
				}
				break;
			case 4:
				{
				setState(88);
				function_call();
				}
				break;
			case 5:
				{
				setState(89);
				table();
				}
				break;
			case 6:
				{
				setState(90);
				table_access();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(93);
					if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
					setState(94);
					operators();
					setState(95);
					expression(5);
					}
					} 
				}
				setState(101);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorsContext extends ParserRuleContext {
		public Comparison_operatorContext comparison_operator() {
			return getRuleContext(Comparison_operatorContext.class,0);
		}
		public Arith_operatorContext arith_operator() {
			return getRuleContext(Arith_operatorContext.class,0);
		}
		public Logical_operatorContext logical_operator() {
			return getRuleContext(Logical_operatorContext.class,0);
		}
		public OperatorsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operators; }
	}

	public final OperatorsContext operators() throws RecognitionException {
		OperatorsContext _localctx = new OperatorsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_operators);
		try {
			setState(105);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				comparison_operator();
				}
				break;
			case T__9:
			case T__10:
			case T__11:
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				arith_operator();
				}
				break;
			case KW_AND:
			case KW_OR:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				logical_operator();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Comparison_operatorContext extends ParserRuleContext {
		public Comparison_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparison_operator; }
	}

	public final Comparison_operatorContext comparison_operator() throws RecognitionException {
		Comparison_operatorContext _localctx = new Comparison_operatorContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_comparison_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1008L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Arith_operatorContext extends ParserRuleContext {
		public Arith_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arith_operator; }
	}

	public final Arith_operatorContext arith_operator() throws RecognitionException {
		Arith_operatorContext _localctx = new Arith_operatorContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_arith_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 15360L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Logical_operatorContext extends ParserRuleContext {
		public TerminalNode KW_AND() { return getToken(lua_grammar_antlr4Parser.KW_AND, 0); }
		public TerminalNode KW_OR() { return getToken(lua_grammar_antlr4Parser.KW_OR, 0); }
		public Logical_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logical_operator; }
	}

	public final Logical_operatorContext logical_operator() throws RecognitionException {
		Logical_operatorContext _localctx = new Logical_operatorContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_logical_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			_la = _input.LA(1);
			if ( !(_la==KW_AND || _la==KW_OR) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CommentContext extends ParserRuleContext {
		public Single_line_commentContext single_line_comment() {
			return getRuleContext(Single_line_commentContext.class,0);
		}
		public Multi_line_commentContext multi_line_comment() {
			return getRuleContext(Multi_line_commentContext.class,0);
		}
		public CommentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comment; }
	}

	public final CommentContext comment() throws RecognitionException {
		CommentContext _localctx = new CommentContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_comment);
		try {
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__22:
				enterOuterAlt(_localctx, 1);
				{
				setState(113);
				single_line_comment();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 2);
				{
				setState(114);
				multi_line_comment();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(lua_grammar_antlr4Parser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(lua_grammar_antlr4Parser.STRING, 0); }
		public TerminalNode KW_TRUE() { return getToken(lua_grammar_antlr4Parser.KW_TRUE, 0); }
		public TerminalNode KW_FALSE() { return getToken(lua_grammar_antlr4Parser.KW_FALSE, 0); }
		public TerminalNode KW_NIL() { return getToken(lua_grammar_antlr4Parser.KW_NIL, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3413468208168960L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_callContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Table_insertContext table_insert() {
			return getRuleContext(Table_insertContext.class,0);
		}
		public TerminalNode KW_PRINT() { return getToken(lua_grammar_antlr4Parser.KW_PRINT, 0); }
		public Function_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call; }
	}

	public final Function_callContext function_call() throws RecognitionException {
		Function_callContext _localctx = new Function_callContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_function_call);
		int _la;
		try {
			setState(139);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				identifier();
				setState(120);
				match(T__1);
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7917068104073220L) != 0)) {
					{
					setState(121);
					expression(0);
					setState(126);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__13) {
						{
						{
						setState(122);
						match(T__13);
						setState(123);
						expression(0);
						}
						}
						setState(128);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(131);
				match(T__2);
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(133);
				table_insert();
				}
				break;
			case KW_PRINT:
				enterOuterAlt(_localctx, 3);
				{
				setState(134);
				match(KW_PRINT);
				setState(135);
				match(T__1);
				setState(136);
				expression(0);
				setState(137);
				match(T__2);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Table_insertContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_insertContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_insert; }
	}

	public final Table_insertContext table_insert() throws RecognitionException {
		Table_insertContext _localctx = new Table_insertContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_table_insert);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(T__14);
			setState(142);
			match(T__1);
			setState(143);
			identifier();
			setState(144);
			match(T__13);
			setState(145);
			expression(0);
			setState(146);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_declarationContext extends ParserRuleContext {
		public TerminalNode KW_FUNCTION() { return getToken(lua_grammar_antlr4Parser.KW_FUNCTION, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public Function_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration; }
	}

	public final Function_declarationContext function_declaration() throws RecognitionException {
		Function_declarationContext _localctx = new Function_declarationContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_function_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(KW_FUNCTION);
			setState(149);
			identifier();
			setState(150);
			match(T__1);
			setState(159);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(151);
				identifier();
				setState(156);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__13) {
					{
					{
					setState(152);
					match(T__13);
					setState(153);
					identifier();
					}
					}
					setState(158);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(161);
			match(T__2);
			setState(162);
			block();
			setState(163);
			match(KW_END);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(165);
				statement();
				}
				}
				setState(168); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 8902574119944196L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_statementContext extends ParserRuleContext {
		public TerminalNode KW_IF() { return getToken(lua_grammar_antlr4Parser.KW_IF, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> KW_THEN() { return getTokens(lua_grammar_antlr4Parser.KW_THEN); }
		public TerminalNode KW_THEN(int i) {
			return getToken(lua_grammar_antlr4Parser.KW_THEN, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public List<TerminalNode> KW_ELSEIF() { return getTokens(lua_grammar_antlr4Parser.KW_ELSEIF); }
		public TerminalNode KW_ELSEIF(int i) {
			return getToken(lua_grammar_antlr4Parser.KW_ELSEIF, i);
		}
		public TerminalNode KW_ELSE() { return getToken(lua_grammar_antlr4Parser.KW_ELSE, 0); }
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(KW_IF);
			setState(171);
			expression(0);
			setState(172);
			match(KW_THEN);
			setState(173);
			block();
			setState(181);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==KW_ELSEIF) {
				{
				{
				setState(174);
				match(KW_ELSEIF);
				setState(175);
				expression(0);
				setState(176);
				match(KW_THEN);
				setState(177);
				block();
				}
				}
				setState(183);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_ELSE) {
				{
				setState(184);
				match(KW_ELSE);
				setState(185);
				block();
				}
			}

			setState(188);
			match(KW_END);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_statementContext extends ParserRuleContext {
		public TerminalNode KW_FOR() { return getToken(lua_grammar_antlr4Parser.KW_FOR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode KW_IN() { return getToken(lua_grammar_antlr4Parser.KW_IN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode KW_DO() { return getToken(lua_grammar_antlr4Parser.KW_DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_for_statement);
		int _la;
		try {
			setState(212);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				match(KW_FOR);
				setState(191);
				identifier();
				setState(192);
				match(KW_IN);
				setState(193);
				expression(0);
				setState(194);
				match(KW_DO);
				setState(195);
				block();
				setState(196);
				match(KW_END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				match(KW_FOR);
				setState(199);
				identifier();
				setState(200);
				match(T__0);
				setState(201);
				expression(0);
				setState(202);
				match(T__13);
				setState(203);
				expression(0);
				setState(206);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__13) {
					{
					setState(204);
					match(T__13);
					setState(205);
					expression(0);
					}
				}

				setState(208);
				match(KW_DO);
				setState(209);
				block();
				setState(210);
				match(KW_END);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_statementContext extends ParserRuleContext {
		public TerminalNode KW_WHILE() { return getToken(lua_grammar_antlr4Parser.KW_WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode KW_DO() { return getToken(lua_grammar_antlr4Parser.KW_DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(KW_WHILE);
			setState(215);
			expression(0);
			setState(216);
			match(KW_DO);
			setState(217);
			block();
			setState(218);
			match(KW_END);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TableContext extends ParserRuleContext {
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public TableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table; }
	}

	public final TableContext table() throws RecognitionException {
		TableContext _localctx = new TableContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_table);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			match(T__15);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7917068104073220L) != 0)) {
				{
				setState(221);
				field();
				setState(226);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__13) {
					{
					{
					setState(222);
					match(T__13);
					setState(223);
					field();
					}
					}
					setState(228);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(231);
			match(T__16);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_field);
		try {
			setState(238);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(233);
				identifier();
				setState(234);
				match(T__0);
				setState(235);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(237);
				expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Table_accessContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_accessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_access; }
	}

	public final Table_accessContext table_access() throws RecognitionException {
		Table_accessContext _localctx = new Table_accessContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_table_access);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(240);
			identifier();
			setState(241);
			match(T__17);
			setState(242);
			expression(0);
			setState(243);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Multi_line_commentContext extends ParserRuleContext {
		public Multi_line_commentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multi_line_comment; }
	}

	public final Multi_line_commentContext multi_line_comment() throws RecognitionException {
		Multi_line_commentContext _localctx = new Multi_line_commentContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_multi_line_comment);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			match(T__19);
			setState(249);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(246);
					_la = _input.LA(1);
					if ( _la <= 0 || (_la==T__20) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					} 
				}
				setState(251);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			setState(252);
			match(T__21);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Single_line_commentContext extends ParserRuleContext {
		public Single_line_commentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_single_line_comment; }
	}

	public final Single_line_commentContext single_line_comment() throws RecognitionException {
		Single_line_commentContext _localctx = new Single_line_commentContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_single_line_comment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(254);
			match(T__22);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Print_statementContext extends ParserRuleContext {
		public TerminalNode KW_PRINT() { return getToken(lua_grammar_antlr4Parser.KW_PRINT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Print_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_statement; }
	}

	public final Print_statementContext print_statement() throws RecognitionException {
		Print_statementContext _localctx = new Print_statementContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_print_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			match(KW_PRINT);
			setState(257);
			match(T__1);
			setState(258);
			expression(0);
			setState(259);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public List<TerminalNode> LETTER() { return getTokens(lua_grammar_antlr4Parser.LETTER); }
		public TerminalNode LETTER(int i) {
			return getToken(lua_grammar_antlr4Parser.LETTER, i);
		}
		public List<TerminalNode> DIGIT() { return getTokens(lua_grammar_antlr4Parser.DIGIT); }
		public TerminalNode DIGIT(int i) {
			return getToken(lua_grammar_antlr4Parser.DIGIT, i);
		}
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_identifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(LETTER);
			setState(265);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(262);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 13510798898888704L) != 0)) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					} 
				}
				setState(267);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 5:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00016\u010d\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0001\u0000\u0001\u0000\u0005\u00007\b\u0000"+
		"\n\u0000\f\u0000:\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0003\u0001A\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002F\b\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0003\u0004"+
		"K\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\\\b\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005b\b\u0005"+
		"\n\u0005\f\u0005e\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006"+
		"j\b\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\n\u0001\n\u0003\nt\b\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0005\f}\b\f\n\f\f\f\u0080\t\f\u0003\f\u0082\b\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u008c"+
		"\b\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e"+
		"\u009b\b\u000e\n\u000e\f\u000e\u009e\t\u000e\u0003\u000e\u00a0\b\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0004\u000f"+
		"\u00a7\b\u000f\u000b\u000f\f\u000f\u00a8\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0005\u0010\u00b4\b\u0010\n\u0010\f\u0010\u00b7\t\u0010\u0001\u0010"+
		"\u0001\u0010\u0003\u0010\u00bb\b\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00cf\b\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00d5\b\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u00e1\b\u0013\n\u0013"+
		"\f\u0013\u00e4\t\u0013\u0003\u0013\u00e6\b\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u00ef\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0016\u0001\u0016\u0005\u0016\u00f8\b\u0016\n\u0016\f\u0016\u00fb"+
		"\t\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0005"+
		"\u0019\u0108\b\u0019\n\u0019\f\u0019\u010b\t\u0019\u0001\u0019\u0000\u0001"+
		"\n\u001a\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018"+
		"\u001a\u001c\u001e \"$&(*,.02\u0000\u0006\u0001\u0000\u0004\t\u0001\u0000"+
		"\n\r\u0002\u0000\u001d\u001d))\u0004\u0000##\'\'--23\u0001\u0000\u0015"+
		"\u0015\u0002\u0000\u0018\u001845\u0114\u00008\u0001\u0000\u0000\u0000"+
		"\u0002@\u0001\u0000\u0000\u0000\u0004E\u0001\u0000\u0000\u0000\u0006G"+
		"\u0001\u0000\u0000\u0000\bJ\u0001\u0000\u0000\u0000\n[\u0001\u0000\u0000"+
		"\u0000\fi\u0001\u0000\u0000\u0000\u000ek\u0001\u0000\u0000\u0000\u0010"+
		"m\u0001\u0000\u0000\u0000\u0012o\u0001\u0000\u0000\u0000\u0014s\u0001"+
		"\u0000\u0000\u0000\u0016u\u0001\u0000\u0000\u0000\u0018\u008b\u0001\u0000"+
		"\u0000\u0000\u001a\u008d\u0001\u0000\u0000\u0000\u001c\u0094\u0001\u0000"+
		"\u0000\u0000\u001e\u00a6\u0001\u0000\u0000\u0000 \u00aa\u0001\u0000\u0000"+
		"\u0000\"\u00d4\u0001\u0000\u0000\u0000$\u00d6\u0001\u0000\u0000\u0000"+
		"&\u00dc\u0001\u0000\u0000\u0000(\u00ee\u0001\u0000\u0000\u0000*\u00f0"+
		"\u0001\u0000\u0000\u0000,\u00f5\u0001\u0000\u0000\u0000.\u00fe\u0001\u0000"+
		"\u0000\u00000\u0100\u0001\u0000\u0000\u00002\u0105\u0001\u0000\u0000\u0000"+
		"47\u0003\u0002\u0001\u000057\u0003\u0014\n\u000064\u0001\u0000\u0000\u0000"+
		"65\u0001\u0000\u0000\u00007:\u0001\u0000\u0000\u000086\u0001\u0000\u0000"+
		"\u000089\u0001\u0000\u0000\u00009\u0001\u0001\u0000\u0000\u0000:8\u0001"+
		"\u0000\u0000\u0000;A\u0003\b\u0004\u0000<A\u0003\n\u0005\u0000=A\u0003"+
		"\u0004\u0002\u0000>A\u0003\u001c\u000e\u0000?A\u0003\u0018\f\u0000@;\u0001"+
		"\u0000\u0000\u0000@<\u0001\u0000\u0000\u0000@=\u0001\u0000\u0000\u0000"+
		"@>\u0001\u0000\u0000\u0000@?\u0001\u0000\u0000\u0000A\u0003\u0001\u0000"+
		"\u0000\u0000BF\u0003 \u0010\u0000CF\u0003\"\u0011\u0000DF\u0003$\u0012"+
		"\u0000EB\u0001\u0000\u0000\u0000EC\u0001\u0000\u0000\u0000ED\u0001\u0000"+
		"\u0000\u0000F\u0005\u0001\u0000\u0000\u0000GH\u0005\u0019\u0000\u0000"+
		"H\u0007\u0001\u0000\u0000\u0000IK\u00050\u0000\u0000JI\u0001\u0000\u0000"+
		"\u0000JK\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LM\u00032\u0019"+
		"\u0000MN\u0005\u0001\u0000\u0000NO\u0003\n\u0005\u0000OP\u0003\u0006\u0003"+
		"\u0000P\t\u0001\u0000\u0000\u0000QR\u0006\u0005\uffff\uffff\u0000R\\\u0003"+
		"\u0016\u000b\u0000S\\\u00032\u0019\u0000TU\u0005\u0002\u0000\u0000UV\u0003"+
		"\n\u0005\u0000VW\u0005\u0003\u0000\u0000W\\\u0001\u0000\u0000\u0000X\\"+
		"\u0003\u0018\f\u0000Y\\\u0003&\u0013\u0000Z\\\u0003*\u0015\u0000[Q\u0001"+
		"\u0000\u0000\u0000[S\u0001\u0000\u0000\u0000[T\u0001\u0000\u0000\u0000"+
		"[X\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000[Z\u0001\u0000\u0000"+
		"\u0000\\c\u0001\u0000\u0000\u0000]^\n\u0004\u0000\u0000^_\u0003\f\u0006"+
		"\u0000_`\u0003\n\u0005\u0005`b\u0001\u0000\u0000\u0000a]\u0001\u0000\u0000"+
		"\u0000be\u0001\u0000\u0000\u0000ca\u0001\u0000\u0000\u0000cd\u0001\u0000"+
		"\u0000\u0000d\u000b\u0001\u0000\u0000\u0000ec\u0001\u0000\u0000\u0000"+
		"fj\u0003\u000e\u0007\u0000gj\u0003\u0010\b\u0000hj\u0003\u0012\t\u0000"+
		"if\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000ih\u0001\u0000\u0000"+
		"\u0000j\r\u0001\u0000\u0000\u0000kl\u0007\u0000\u0000\u0000l\u000f\u0001"+
		"\u0000\u0000\u0000mn\u0007\u0001\u0000\u0000n\u0011\u0001\u0000\u0000"+
		"\u0000op\u0007\u0002\u0000\u0000p\u0013\u0001\u0000\u0000\u0000qt\u0003"+
		".\u0017\u0000rt\u0003,\u0016\u0000sq\u0001\u0000\u0000\u0000sr\u0001\u0000"+
		"\u0000\u0000t\u0015\u0001\u0000\u0000\u0000uv\u0007\u0003\u0000\u0000"+
		"v\u0017\u0001\u0000\u0000\u0000wx\u00032\u0019\u0000x\u0081\u0005\u0002"+
		"\u0000\u0000y~\u0003\n\u0005\u0000z{\u0005\u000e\u0000\u0000{}\u0003\n"+
		"\u0005\u0000|z\u0001\u0000\u0000\u0000}\u0080\u0001\u0000\u0000\u0000"+
		"~|\u0001\u0000\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0082"+
		"\u0001\u0000\u0000\u0000\u0080~\u0001\u0000\u0000\u0000\u0081y\u0001\u0000"+
		"\u0000\u0000\u0081\u0082\u0001\u0000\u0000\u0000\u0082\u0083\u0001\u0000"+
		"\u0000\u0000\u0083\u0084\u0005\u0003\u0000\u0000\u0084\u008c\u0001\u0000"+
		"\u0000\u0000\u0085\u008c\u0003\u001a\r\u0000\u0086\u0087\u0005\u001c\u0000"+
		"\u0000\u0087\u0088\u0005\u0002\u0000\u0000\u0088\u0089\u0003\n\u0005\u0000"+
		"\u0089\u008a\u0005\u0003\u0000\u0000\u008a\u008c\u0001\u0000\u0000\u0000"+
		"\u008bw\u0001\u0000\u0000\u0000\u008b\u0085\u0001\u0000\u0000\u0000\u008b"+
		"\u0086\u0001\u0000\u0000\u0000\u008c\u0019\u0001\u0000\u0000\u0000\u008d"+
		"\u008e\u0005\u000f\u0000\u0000\u008e\u008f\u0005\u0002\u0000\u0000\u008f"+
		"\u0090\u00032\u0019\u0000\u0090\u0091\u0005\u000e\u0000\u0000\u0091\u0092"+
		"\u0003\n\u0005\u0000\u0092\u0093\u0005\u0003\u0000\u0000\u0093\u001b\u0001"+
		"\u0000\u0000\u0000\u0094\u0095\u00051\u0000\u0000\u0095\u0096\u00032\u0019"+
		"\u0000\u0096\u009f\u0005\u0002\u0000\u0000\u0097\u009c\u00032\u0019\u0000"+
		"\u0098\u0099\u0005\u000e\u0000\u0000\u0099\u009b\u00032\u0019\u0000\u009a"+
		"\u0098\u0001\u0000\u0000\u0000\u009b\u009e\u0001\u0000\u0000\u0000\u009c"+
		"\u009a\u0001\u0000\u0000\u0000\u009c\u009d\u0001\u0000\u0000\u0000\u009d"+
		"\u00a0\u0001\u0000\u0000\u0000\u009e\u009c\u0001\u0000\u0000\u0000\u009f"+
		"\u0097\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000\u0000\u0000\u00a0"+
		"\u00a1\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\u0003\u0000\u0000\u00a2"+
		"\u00a3\u0003\u001e\u000f\u0000\u00a3\u00a4\u0005\"\u0000\u0000\u00a4\u001d"+
		"\u0001\u0000\u0000\u0000\u00a5\u00a7\u0003\u0002\u0001\u0000\u00a6\u00a5"+
		"\u0001\u0000\u0000\u0000\u00a7\u00a8\u0001\u0000\u0000\u0000\u00a8\u00a6"+
		"\u0001\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000\u0000\u00a9\u001f"+
		"\u0001\u0000\u0000\u0000\u00aa\u00ab\u0005&\u0000\u0000\u00ab\u00ac\u0003"+
		"\n\u0005\u0000\u00ac\u00ad\u0005,\u0000\u0000\u00ad\u00b5\u0003\u001e"+
		"\u000f\u0000\u00ae\u00af\u0005!\u0000\u0000\u00af\u00b0\u0003\n\u0005"+
		"\u0000\u00b0\u00b1\u0005,\u0000\u0000\u00b1\u00b2\u0003\u001e\u000f\u0000"+
		"\u00b2\u00b4\u0001\u0000\u0000\u0000\u00b3\u00ae\u0001\u0000\u0000\u0000"+
		"\u00b4\u00b7\u0001\u0000\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000"+
		"\u00b5\u00b6\u0001\u0000\u0000\u0000\u00b6\u00ba\u0001\u0000\u0000\u0000"+
		"\u00b7\u00b5\u0001\u0000\u0000\u0000\u00b8\u00b9\u0005 \u0000\u0000\u00b9"+
		"\u00bb\u0003\u001e\u000f\u0000\u00ba\u00b8\u0001\u0000\u0000\u0000\u00ba"+
		"\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc"+
		"\u00bd\u0005\"\u0000\u0000\u00bd!\u0001\u0000\u0000\u0000\u00be\u00bf"+
		"\u0005$\u0000\u0000\u00bf\u00c0\u00032\u0019\u0000\u00c0\u00c1\u0005\u001b"+
		"\u0000\u0000\u00c1\u00c2\u0003\n\u0005\u0000\u00c2\u00c3\u0005\u001f\u0000"+
		"\u0000\u00c3\u00c4\u0003\u001e\u000f\u0000\u00c4\u00c5\u0005\"\u0000\u0000"+
		"\u00c5\u00d5\u0001\u0000\u0000\u0000\u00c6\u00c7\u0005$\u0000\u0000\u00c7"+
		"\u00c8\u00032\u0019\u0000\u00c8\u00c9\u0005\u0001\u0000\u0000\u00c9\u00ca"+
		"\u0003\n\u0005\u0000\u00ca\u00cb\u0005\u000e\u0000\u0000\u00cb\u00ce\u0003"+
		"\n\u0005\u0000\u00cc\u00cd\u0005\u000e\u0000\u0000\u00cd\u00cf\u0003\n"+
		"\u0005\u0000\u00ce\u00cc\u0001\u0000\u0000\u0000\u00ce\u00cf\u0001\u0000"+
		"\u0000\u0000\u00cf\u00d0\u0001\u0000\u0000\u0000\u00d0\u00d1\u0005\u001f"+
		"\u0000\u0000\u00d1\u00d2\u0003\u001e\u000f\u0000\u00d2\u00d3\u0005\"\u0000"+
		"\u0000\u00d3\u00d5\u0001\u0000\u0000\u0000\u00d4\u00be\u0001\u0000\u0000"+
		"\u0000\u00d4\u00c6\u0001\u0000\u0000\u0000\u00d5#\u0001\u0000\u0000\u0000"+
		"\u00d6\u00d7\u0005/\u0000\u0000\u00d7\u00d8\u0003\n\u0005\u0000\u00d8"+
		"\u00d9\u0005\u001f\u0000\u0000\u00d9\u00da\u0003\u001e\u000f\u0000\u00da"+
		"\u00db\u0005\"\u0000\u0000\u00db%\u0001\u0000\u0000\u0000\u00dc\u00e5"+
		"\u0005\u0010\u0000\u0000\u00dd\u00e2\u0003(\u0014\u0000\u00de\u00df\u0005"+
		"\u000e\u0000\u0000\u00df\u00e1\u0003(\u0014\u0000\u00e0\u00de\u0001\u0000"+
		"\u0000\u0000\u00e1\u00e4\u0001\u0000\u0000\u0000\u00e2\u00e0\u0001\u0000"+
		"\u0000\u0000\u00e2\u00e3\u0001\u0000\u0000\u0000\u00e3\u00e6\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e5\u00dd\u0001\u0000"+
		"\u0000\u0000\u00e5\u00e6\u0001\u0000\u0000\u0000\u00e6\u00e7\u0001\u0000"+
		"\u0000\u0000\u00e7\u00e8\u0005\u0011\u0000\u0000\u00e8\'\u0001\u0000\u0000"+
		"\u0000\u00e9\u00ea\u00032\u0019\u0000\u00ea\u00eb\u0005\u0001\u0000\u0000"+
		"\u00eb\u00ec\u0003\n\u0005\u0000\u00ec\u00ef\u0001\u0000\u0000\u0000\u00ed"+
		"\u00ef\u0003\n\u0005\u0000\u00ee\u00e9\u0001\u0000\u0000\u0000\u00ee\u00ed"+
		"\u0001\u0000\u0000\u0000\u00ef)\u0001\u0000\u0000\u0000\u00f0\u00f1\u0003"+
		"2\u0019\u0000\u00f1\u00f2\u0005\u0012\u0000\u0000\u00f2\u00f3\u0003\n"+
		"\u0005\u0000\u00f3\u00f4\u0005\u0013\u0000\u0000\u00f4+\u0001\u0000\u0000"+
		"\u0000\u00f5\u00f9\u0005\u0014\u0000\u0000\u00f6\u00f8\b\u0004\u0000\u0000"+
		"\u00f7\u00f6\u0001\u0000\u0000\u0000\u00f8\u00fb\u0001\u0000\u0000\u0000"+
		"\u00f9\u00f7\u0001\u0000\u0000\u0000\u00f9\u00fa\u0001\u0000\u0000\u0000"+
		"\u00fa\u00fc\u0001\u0000\u0000\u0000\u00fb\u00f9\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fd\u0005\u0016\u0000\u0000\u00fd-\u0001\u0000\u0000\u0000\u00fe"+
		"\u00ff\u0005\u0017\u0000\u0000\u00ff/\u0001\u0000\u0000\u0000\u0100\u0101"+
		"\u0005\u001c\u0000\u0000\u0101\u0102\u0005\u0002\u0000\u0000\u0102\u0103"+
		"\u0003\n\u0005\u0000\u0103\u0104\u0005\u0003\u0000\u0000\u01041\u0001"+
		"\u0000\u0000\u0000\u0105\u0109\u00054\u0000\u0000\u0106\u0108\u0007\u0005"+
		"\u0000\u0000\u0107\u0106\u0001\u0000\u0000\u0000\u0108\u010b\u0001\u0000"+
		"\u0000\u0000\u0109\u0107\u0001\u0000\u0000\u0000\u0109\u010a\u0001\u0000"+
		"\u0000\u0000\u010a3\u0001\u0000\u0000\u0000\u010b\u0109\u0001\u0000\u0000"+
		"\u0000\u001868@EJ[cis~\u0081\u008b\u009c\u009f\u00a8\u00b5\u00ba\u00ce"+
		"\u00d4\u00e2\u00e5\u00ee\u00f9\u0109";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}